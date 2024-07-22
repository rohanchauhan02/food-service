package repository

import (
	"github.com/labstack/echo"
	"github.com/rohanchauhan02/food-service/models"
)

// FetchBalanceSheet is a method to get the balance sheet of a bussiness
func (r *repoHandler) FetchBalanceSheet(c echo.Context, data models.BussinessDetail) ([]models.BalanceSheet, error) {
	resp := []models.BalanceSheet{
		{
			Year:         2020,
			Month:        12,
			ProfitOrLoss: 250000,
			AssetsValue:  1234,
		},
		{
			Year:         2020,
			Month:        11,
			ProfitOrLoss: 1150,
			AssetsValue:  5789,
		},
		{
			Year:         2020,
			Month:        10,
			ProfitOrLoss: 2500,
			AssetsValue:  22345,
		},
		{
			Year:         2020,
			Month:        9,
			ProfitOrLoss: -187000,
			AssetsValue:  223452,
		},
	}
	return resp, nil
}

// PreAssement method is used to calculate the loan amont on the basis of company balance sheet
func (r *repoHandler) PreAssessment(balanceSheet []models.BalanceSheet, loanAmount int) float64 {
	profitInLast12Months := false
	averageAssetValue := 0.0

	for _, entry := range balanceSheet {
		if entry.ProfitOrLoss > 0 {
			profitInLast12Months = true
		}
		averageAssetValue += entry.AssetsValue
	}
	averageAssetValue /= float64(len(balanceSheet))

	if profitInLast12Months {
		return float64(loanAmount) * 0.6
	} else if averageAssetValue > float64(loanAmount) {
		return float64(loanAmount) * 1
	} else {
		return float64(loanAmount) * 0.2
	}
}
