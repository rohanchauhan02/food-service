package usecase

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rohanchauhan02/food-service/mocks/mock_bank"
	"github.com/rohanchauhan02/food-service/models"
)

func Test_usecaseHandler_LoanEnquiry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockBank := mock_bank.NewMockRepository(mockCtrl)

	data := []models.BalanceSheet{
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
	businessDetail := models.BussinessDetail{
		BussinessName: "WEBX",
		BussinessGST:  "hjj",
	}
	tests := []struct {
		name     string
		prepFunc func(*mock_bank.MockRepository)
		payload  *models.LoanEnquiryRequest
		want     *models.LoanEnquiryResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "TEST 1",
			prepFunc: func(mock_bank *mock_bank.MockRepository) {
				mock_bank.EXPECT().FetchBalanceSheet(gomock.Any(), businessDetail).Return(data, nil).Times(1)
				mock_bank.EXPECT().PreAssessment(gomock.Any(), 100).Return(60.0)
			},
			payload: &models.LoanEnquiryRequest{
				BussinessName: "WEBX",
				BussinessGST:  "hjj",
				ContactNumber: 1234,
				Provider:      "ghh",
				LoanAmount:    100,
			},
			want: &models.LoanEnquiryResponse{
				Name:          "WEBX",
				PreAssessment: 60,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecaseHandler{
				repository: mockBank,
			}
			if tt.prepFunc != nil {
				tt.prepFunc(mockBank)
			}
			got, err := u.LoanEnquiry(nil, tt.payload)

			if (err != nil) != tt.wantErr {
				t.Errorf("usecaseHandler.LoanEnquiry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecaseHandler.LoanEnquiry() = %v, want %v", got, tt.want)
			}
		})
	}
}
