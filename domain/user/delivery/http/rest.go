package http

import (
	"github.com/labstack/echo"
	"github.com/rohanchauhan02/food-service/domain/user"
	"github.com/rohanchauhan02/food-service/models"
)

type userHandler struct {
	usecase user.Usecase
}

func NewUserHandler(e *echo.Echo, usecase user.Usecase) {
	handler := userHandler{
		usecase: usecase,
	}
	e.POST("/api/signup", handler.Register)
	e.POST("/api/login", handler.Login)
}
func (h *userHandler) Register(c echo.Context) error {
	reqPayload := &models.UserRegisterRequest{}
	if err := c.Bind(reqPayload); err != nil {
		return err
	}
	if err := c.Validate(reqPayload); err != nil {
		return err
	}
	resp, err := h.usecase.Register(c, *reqPayload)
	if err != nil {
		return err
	}

	return c.JSON(200, resp)
}

func (h *userHandler) Login(c echo.Context) error {
	reqPayload := &models.UserLogin{}
	if err := c.Bind(reqPayload); err != nil {
		return err
	}
	if err := c.Validate(reqPayload); err != nil {
		return err
	}
	resp, err := h.usecase.Login(c, *reqPayload)
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"token": resp,
		"msg":   "Successfully logged in",
	}
	return c.JSON(200, data)
}
