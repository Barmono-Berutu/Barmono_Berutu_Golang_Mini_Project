package service

import (
	"project_mini_golang/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserID int    `json:"user_id"` // Perbaikan tag JSON
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateJWT(email string, userID int) (string, error)
}

type jwtService struct {
	config *config.JWTConfig
}

func NewJWTService(cfg *config.JWTConfig) JWTService {
	return &jwtService{config: cfg}
}

func (s *jwtService) GenerateJWT(email string, id int) (string, error) {
	claims := &JwtCustomClaims{
		UserID: id,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}
	// Membuat token dengan claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.config.SecretKey))

	if err != nil {
		return "", err
	}
	return t, nil
}
