package healthcheck

import (
	"github.com/labstack/echo"
)

type Repository interface {
	MysqlHealthCheck() (bool, error)
}

type Usecase interface {
	HealthCheck(c echo.Context) (bool, error)
}
