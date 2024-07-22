package middleware

import (
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/rohanchauhan02/food-service/models"
	"github.com/rohanchauhan02/food-service/shared/util"
)

func JWTAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.(*util.CustomApplicationContext)
		conf := ac.SharedConfig

		tokenHeader := c.Request().Header.Get("Authorization")
		if tokenHeader == "" {
			return c.JSON(403, map[string]string{
				"message": "Missing auth token",
			})
		}

		splitedTokenHeader := strings.Split(tokenHeader, " ")
		if len(splitedTokenHeader) != 2 {
			return c.JSON(403, map[string]string{
				"message": "Invalid/Malformed auth token",
			})
		}

		tokenPart := splitedTokenHeader[1]

		tk := &models.UserJWT{}
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.GetJWTSecret()), nil
		})
		if err != nil {
			return c.JSON(403, map[string]string{
				"message": "Malformed authentication token",
				"error":   err.Error(),
			})
		}
		if !token.Valid {
			return c.JSON(403, map[string]string{
				"message": "Token is not valid.",
			})
		}
		ac.UserJWT = tk
		return next(c)
	}
}
