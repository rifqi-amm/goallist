package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)


func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bearer token-blablala=
		authorizationFromHeader := c.Request().Header.Get("authorization")

		// Replace Bearer, Get token-blablala=
		tokenString := strings.ReplaceAll(authorizationFromHeader, "Bearer ", "")

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("asdasdasd"), nil
		})
		if err != nil {
			return c.String(http.StatusForbidden, "token salah")
		}

		c.Set("email", claims["userId"])
		return next(c)
	}
}