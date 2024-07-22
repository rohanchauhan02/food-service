package user

import (
	"github.com/labstack/echo"
	"github.com/rohanchauhan02/food-service/models"
)

type Usecase interface {
	Register(c echo.Context, payload models.UserRegisterRequest) (string, error)
	Login(c echo.Context, payload models.UserLogin) (string, error)
}
type Repository interface {
	Register(dto models.User) error
	Login(dto models.UserLogin) (models.User, error)
}
