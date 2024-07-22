package http

import (
	"log"

	"github.com/labstack/echo"
	"github.com/rohanchauhan02/food-service/domain/healthcheck"
	"github.com/rohanchauhan02/food-service/shared/util"
)

type handlerHealthcheck struct {
	usecase healthcheck.Usecase
}

func NewHandlerHealthcheck(e *echo.Echo, usecase healthcheck.Usecase) {
	handler := &handlerHealthcheck{
		usecase: usecase,
	}
	e.GET("/api/healthz", handler.HealthCheck)
}

func (h *handlerHealthcheck) HealthCheck(c echo.Context) error {
	ac := c.(*util.CustomApplicationContext)
	ac.Logger().Info("Healthcheck endpoint hit")
	res, err := h.usecase.HealthCheck(c)
	if err != nil {
		log.Fatalf("[FAILED] failed to get healthcheck response: %s", err.Error())
		return err
	}
	ac.Logger().Info("[SUCCESS] healthcheck response :%s", res)
	return nil
}
