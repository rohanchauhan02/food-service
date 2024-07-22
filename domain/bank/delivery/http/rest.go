package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rohanchauhan02/food-service/domain/bank"
	"github.com/rohanchauhan02/food-service/models"
	"github.com/rohanchauhan02/food-service/shared/middleware"
)

type handlerBank struct {
	usecase bank.Usecase
}

func NewHandlerBank(e *echo.Echo, usecase bank.Usecase) {
	handler := &handlerBank{
		usecase: usecase,
	}
	e.POST("/api/loan-enquiry", handler.LoanEnquiry, middleware.JWTAuthentication)
	e.POST("/api/quote", handler.CreateQuote, middleware.JWTAuthentication)
}

func (h *handlerBank) LoanEnquiry(c echo.Context) error {
	reqPayload := &models.LoanEnquiryRequest{}

	if err := c.Bind(&reqPayload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(reqPayload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := h.usecase.LoanEnquiry(c, reqPayload)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *handlerBank) CreateQuote(c echo.Context) error {
	return nil
}
