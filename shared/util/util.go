package util

import (
	"fmt"
	"time"

	"github.com/fgrosse/goldi"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/rohanchauhan02/food-service/models"
	"github.com/rohanchauhan02/food-service/shared/config"
)

type CustomApplicationContext struct {
	echo.Context
	Container    *goldi.Container
	SharedConfig config.ImmutableConfigInterface
	MysqlSession *gorm.DB
	UserJWT      *models.UserJWT
}

type CutomValidator struct {
	Validator *validator.Validate
}

func (c *CutomValidator) Validate(i interface{}) error {
	return c.Validator.Struct(i)
}

func DefaultValidator() *CutomValidator {
	return &CutomValidator{Validator: validator.New()}
}

func TokenCreater(c echo.Context, data models.User) (string, error) {
	ac := c.(*CustomApplicationContext)
	conf := ac.SharedConfig
	secretKey := []byte(conf.GetJWTSecret())

	// Create the claims
	claims := &models.UserJWT{
		Email:    data.Email,
		UserID:   data.ID,
		UserName: data.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "food-service",
		},
	}

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error creating token:", err)
		return "", err
	}

	return tokenString, nil
}
