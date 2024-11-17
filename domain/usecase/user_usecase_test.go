package usecase

import (
	"project_mini_golang/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepoDummy struct{}

func (ard AuthRepoDummy) GetByUsername(email string) (*model.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	return &model.User{
		ID:       1,
		Email:    email,
		Password: string(hashedPassword),
	}, nil
}

func (ard AuthRepoDummy) CreateUser(user *model.User) error {
	return nil
}

type JWTServiceDummy struct{}

func (j JWTServiceDummy) GenerateJWT(email string, userID int) (string, error) {
	return "dummy_token", nil
}

var AuthUsecaseTest UserUsecase

func setup() {
	authRepoDummy := AuthRepoDummy{}
	jwtServiceDummy := JWTServiceDummy{}
	AuthUsecaseTest = NewAuthUsecase(authRepoDummy, jwtServiceDummy)
}

func TestAuthService_Login(t *testing.T) {
	setup()

	t.Run("sukses login", func(t *testing.T) {
		token, err := AuthUsecaseTest.Login("octo@gmail.com", "123456")
		assert.Nil(t, err)
		assert.Equal(t, "dummy_token", token)
	})

	t.Run("gagal login email kosong", func(t *testing.T) {
		token, err := AuthUsecaseTest.Login("", "123456")
		assert.NotNil(t, err)
		assert.Equal(t, "email is required", err.Error())
		assert.Equal(t, "", token)
	})

	t.Run("gagal login password salah", func(t *testing.T) {
		token, err := AuthUsecaseTest.Login("octo@gmail.com", "wrongpassword")
		assert.NotNil(t, err)
		assert.Equal(t, "invalid credentials", err.Error())
		assert.Equal(t, "", token)
	})
}

func TestAuthService_Register(t *testing.T) {
	setup()

	t.Run("sukses register", func(t *testing.T) {
		user := &model.User{
			Email:    "octo@gmail.com",
			Password: "123456",
		}

		err := AuthUsecaseTest.Register(user)

		assert.Nil(t, err)
	})

	t.Run("gagal register email kosong", func(t *testing.T) {
		user := &model.User{
			Email:    "",
			Password: "123456",
		}

		err := AuthUsecaseTest.Register(user)

		assert.NotNil(t, err)
		assert.Equal(t, "email is required", err.Error())
	})

	t.Run("gagal register password kosong", func(t *testing.T) {
		user := &model.User{
			Email:    "octo@gmail.com",
			Password: "",
		}

		err := AuthUsecaseTest.Register(user)

		assert.NotNil(t, err)
		assert.Equal(t, "password is required", err.Error())
	})
}
