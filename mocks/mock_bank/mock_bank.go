// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/bank/bank.go

// Package mock_bank is a generated GoMock package.
package mock_bank

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo"
	models "github.com/rohanchauhan02/food-service/models"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// LoanEnquiry mocks base method.
func (m *MockUsecase) LoanEnquiry(c echo.Context, payload *models.LoanEnquiryRequest) (*models.LoanEnquiryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoanEnquiry", c, payload)
	ret0, _ := ret[0].(*models.LoanEnquiryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoanEnquiry indicates an expected call of LoanEnquiry.
func (mr *MockUsecaseMockRecorder) LoanEnquiry(c, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoanEnquiry", reflect.TypeOf((*MockUsecase)(nil).LoanEnquiry), c, payload)
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateLead mocks base method.
func (m *MockRepository) CreateLead() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLead")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLead indicates an expected call of CreateLead.
func (mr *MockRepositoryMockRecorder) CreateLead() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLead", reflect.TypeOf((*MockRepository)(nil).CreateLead))
}

// FetchBalanceSheet mocks base method.
func (m *MockRepository) FetchBalanceSheet(c echo.Context, data models.BussinessDetail) ([]models.BalanceSheet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBalanceSheet", c, data)
	ret0, _ := ret[0].([]models.BalanceSheet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBalanceSheet indicates an expected call of FetchBalanceSheet.
func (mr *MockRepositoryMockRecorder) FetchBalanceSheet(c, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBalanceSheet", reflect.TypeOf((*MockRepository)(nil).FetchBalanceSheet), c, data)
}

// PreAssessment mocks base method.
func (m *MockRepository) PreAssessment(balanceSheet []models.BalanceSheet, loanAmount int) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreAssessment", balanceSheet, loanAmount)
	ret0, _ := ret[0].(float64)
	return ret0
}

// PreAssessment indicates an expected call of PreAssessment.
func (mr *MockRepositoryMockRecorder) PreAssessment(balanceSheet, loanAmount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreAssessment", reflect.TypeOf((*MockRepository)(nil).PreAssessment), balanceSheet, loanAmount)
}
