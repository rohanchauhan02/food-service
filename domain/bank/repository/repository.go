package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rohanchauhan02/food-service/domain/bank"
)

type repoHandler struct {
	mysqlSession *gorm.DB
}

func NewBankRepository(mysqlSession *gorm.DB) bank.Repository {
	return &repoHandler{
		mysqlSession: mysqlSession,
	}
}

func (r *repoHandler) CreateLead() error {
	return nil
}
