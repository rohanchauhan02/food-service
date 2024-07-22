package usecase

import (
	"github.com/labstack/echo"
	"github.com/rohanchauhan02/food-service/domain/bank"
	"github.com/rohanchauhan02/food-service/models"
)

type usecaseHandler struct {
	repository bank.Repository
}

func NewBankUsecase(repository bank.Repository) bank.Usecase {
	return &usecaseHandler{
		repository: repository,
	}
}

func (u *usecaseHandler) LoanEnquiry(c echo.Context, payload *models.LoanEnquiryRequest) (*models.LoanEnquiryResponse, error) {

	// Create Lead in go routine
	// _ = u.repository.CreateLead()

	// fetch balance sheet
	businessDetails := models.BussinessDetail{
		BussinessName: payload.BussinessName,
		BussinessGST:  payload.BussinessGST,
	}
	balanceSheet, _ := u.repository.FetchBalanceSheet(c, businessDetails)

	// PreAssisment
	preAssesment := u.repository.PreAssessment(balanceSheet, payload.LoanAmount)

	data := &models.LoanEnquiryResponse{
		Name:          payload.BussinessName,
		PreAssessment: float64(preAssesment),
	}

	return data, nil
}
