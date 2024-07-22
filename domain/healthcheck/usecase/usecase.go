package usecase

import (
	"runtime/debug"

	"github.com/labstack/echo"
	"github.com/rohanchauhan02/food-service/domain/healthcheck"
)

type usecase struct {
	repository healthcheck.Repository
}

func NewHelthCheckUsecase(repository healthcheck.Repository) healthcheck.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) HealthCheck(c echo.Context) (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			c.Logger().Errorf("Panic occured :", err)
			debug.PrintStack()
		}
	}()
	_, err := u.repository.MysqlHealthCheck()
	if err != nil {
		return false, err
	}

	return true, nil
}
