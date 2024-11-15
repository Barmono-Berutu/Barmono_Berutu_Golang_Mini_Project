package middleware

import (
	"project_mini_golang/config"
	"project_mini_golang/helper"
	"project_mini_golang/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JWTMiddleware struct {
	config *config.JWTConfig
}

func NewJWTMiddleware(cfg *config.JWTConfig) *JWTMiddleware {
	return &JWTMiddleware{config: cfg}
}

func (m *JWTMiddleware) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h, err := c.Cookie("token")
		if err != nil {
			return helper.JSONErrorResponse(c, 401, "gagal login token tidak ditemukan")

		}

		tokenString := h.Value
		claims := &service.JwtCustomClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(m.config.SecretKey), nil
		})

		if err != nil || !token.Valid {
			return helper.JSONErrorResponse(c, 401, "invalid token")
		}

		c.Set("user_id", claims.UserID)

		// Melanjutkan ke handler berikutnya jika token valid
		return next(c)
	}
}
