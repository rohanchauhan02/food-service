package bank

import (
	"github.com/labstack/echo"
	"github.com/rohanchauhan02/food-service/models"
)

type Usecase interface {
	LoanEnquiry(c echo.Context, payload *models.LoanEnquiryRequest) (*models.LoanEnquiryResponse, error)
}
type Repository interface {

	// DB transection
	CreateLead() error
	// 3rd Party Api's
	FetchBalanceSheet(c echo.Context, data models.BussinessDetail) ([]models.BalanceSheet, error)
	PreAssessment(balanceSheet []models.BalanceSheet, loanAmount int) float64
}
