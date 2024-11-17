package service

import (
	"project_mini_golang/config"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func setupJWTService(secretKey string) JWTService {
	mockConfig := &config.JWTConfig{
		SecretKey: secretKey,
	}
	return NewJWTService(mockConfig)
}

func verifyToken(t *testing.T, token string, secretKey string, expectedEmail string, expectedUserID int) {
	parsedToken, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	assert.Nil(t, err, "Tidak boleh ada error saat parsing token")
	assert.NotNil(t, parsedToken, "Token yang di-parse tidak boleh kosong")

	if claims, ok := parsedToken.Claims.(*JwtCustomClaims); ok && parsedToken.Valid {
		assert.Equal(t, expectedUserID, claims.UserID, "UserID harus sesuai")
		assert.Equal(t, expectedEmail, claims.Email, "Email harus sesuai")
		assert.WithinDuration(t, time.Now().Add(72*time.Hour), claims.ExpiresAt.Time, time.Minute, "Tanggal kedaluwarsa harus dalam rentang yang sesuai")
	} else {
		t.Errorf("Token tidak valid atau tidak mengandung klaim yang diharapkan")
	}
}
func TestGenerateJWT(t *testing.T) {
	t.Run("Berhasil menghasilkan token JWT", func(t *testing.T) {
		jwtService := setupJWTService("testsecretkey")

		email := "octo@gmail.com"
		userID := 1

		token, err := jwtService.GenerateJWT(email, userID)

		assert.Nil(t, err, "Tidak boleh ada error saat menghasilkan token")
		assert.NotEmpty(t, token, "Token tidak boleh kosong")

		verifyToken(t, token, "testsecretkey", email, userID)
	})

	t.Run("Gagal menghasilkan token JWT dengan secret key kosong", func(t *testing.T) {
		jwtService := setupJWTService("")

		email := "octo@gmail.com"
		userID := 1

		token, err := jwtService.GenerateJWT(email, userID)

		assert.NotNil(t, err, "Error harus muncul jika secret key kosong")
		assert.Equal(t, "", token, "Token harus kosong jika secret key kosong")
		assert.Equal(t, "secret key is required", err.Error(), "Pesan error harus sesuai")
	})
}
