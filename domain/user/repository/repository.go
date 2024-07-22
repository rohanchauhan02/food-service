package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rohanchauhan02/food-service/domain/user"
	"github.com/rohanchauhan02/food-service/models"
)

type repository struct {
	mysqlSess *gorm.DB
}

func NewUserRepository(mysqlSess *gorm.DB) user.Repository {
	return &repository{mysqlSess: mysqlSess}
}

func (r *repository) Register(dto models.User) error {
	db := r.mysqlSess.Create(&dto)
	return db.Error
}
func (r *repository) Login(dto models.UserLogin) (models.User, error) {
	var user models.User
	db := r.mysqlSess.Where("email = ? AND password = ?", dto.Email, dto.Password).First(&user)
	return user, db.Error
}
