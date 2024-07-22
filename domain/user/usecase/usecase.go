package usecase

import (
	"github.com/labstack/echo"
	"github.com/rohanchauhan02/food-service/domain/user"
	"github.com/rohanchauhan02/food-service/models"
	"github.com/rohanchauhan02/food-service/shared/util"
)

type usecase struct {
	repository user.Repository
}

func NewUserUsecase(repository user.Repository) user.Usecase {
	return &usecase{repository: repository}
}

func (u *usecase) Register(c echo.Context, payload models.UserRegisterRequest) (string, error) {

	dto := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}
	err := u.repository.Register(dto)
	if err != nil {
		return "", err
	}

	return "Successfully registered", nil
}

func (u *usecase) Login(c echo.Context, payload models.UserLogin) (string, error) {

	resp, err := u.repository.Login(payload)
	if err != nil {
		return "", err
	}

	token, err := util.TokenCreater(c, resp)
	if err != nil {
		return "", err
	}

	return token, nil
}
