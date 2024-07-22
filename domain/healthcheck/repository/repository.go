package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rohanchauhan02/food-service/domain/healthcheck"
)

type repository struct {
	mysqlsess *gorm.DB
}

func NewHealthCheckRepository(mysqlsess *gorm.DB) healthcheck.Repository {
	return &repository{
		mysqlsess: mysqlsess,
	}
}

func (r *repository) MysqlHealthCheck() (bool, error) {
	err := r.mysqlsess.DB().Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}
