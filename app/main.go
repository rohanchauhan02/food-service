package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	bankHandler "github.com/rohanchauhan02/food-service/domain/bank/delivery/http"
	bankRepository "github.com/rohanchauhan02/food-service/domain/bank/repository"
	bankUsecase "github.com/rohanchauhan02/food-service/domain/bank/usecase"

	healthzHandler "github.com/rohanchauhan02/food-service/domain/healthcheck/delivery/http"
	healthzRepository "github.com/rohanchauhan02/food-service/domain/healthcheck/repository"
	healthzUsecase "github.com/rohanchauhan02/food-service/domain/healthcheck/usecase"

	userHandler "github.com/rohanchauhan02/food-service/domain/user/delivery/http"
	userRepository "github.com/rohanchauhan02/food-service/domain/user/repository"
	userUsecase "github.com/rohanchauhan02/food-service/domain/user/usecase"

	"github.com/rohanchauhan02/food-service/shared/config"
	"github.com/rohanchauhan02/food-service/shared/container"
	"github.com/rohanchauhan02/food-service/shared/database"
	"github.com/rohanchauhan02/food-service/shared/util"
)

func main() {
	e := echo.New()
	container := container.DefaultContainer()
	conf := container.MustGet("shared.config").(config.ImmutableConfigInterface)
	mysql := container.MustGet("shared.database").(database.MysqlInterface)
	mysqlSession, err := mysql.OpenMysqlConn()
	if err != nil {
		msgError := fmt.Sprintf("Failed to open mysql connection: %s", err.Error())
		log.Error(msgError)
		panic(err)
	}
	defer mysqlSession.Close()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Validator = util.DefaultValidator()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := &util.CustomApplicationContext{
				Context:      c,
				Container:    container,
				SharedConfig: conf,
				MysqlSession: mysqlSession,
			}
			return next(ac)
		}
	})
	healthCheckRepo := healthzRepository.NewHealthCheckRepository(mysqlSession)
	healthCheckUsecase := healthzUsecase.NewHelthCheckUsecase(healthCheckRepo)
	healthzHandler.NewHandlerHealthcheck(e, healthCheckUsecase)

	userRepo := userRepository.NewUserRepository(mysqlSession)
	userUcase := userUsecase.NewUserUsecase(userRepo)
	userHandler.NewUserHandler(e, userUcase)

	bankRepo := bankRepository.NewBankRepository(mysqlSession)
	bankUse := bankUsecase.NewBankUsecase(bankRepo)
	bankHandler.NewHandlerBank(e, bankUse)

	e.Logger.Info(e.Start(fmt.Sprintf(":%s", conf.GetPort())))
}
