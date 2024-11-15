package controllers

import (
	"net/http"
	"project_mini_golang/domain/model"
	"project_mini_golang/domain/usecase"
	"project_mini_golang/helper"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthUsecase usecase.UserUsecase
}

func NewAuthController(authUsecase usecase.UserUsecase) *AuthController {
	return &AuthController{
		AuthUsecase: authUsecase,
	}
}

func (c *AuthController) Register(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusBadRequest, "gagal mendapatkan data: "+err.Error())
	}

	err := c.AuthUsecase.Register(&user)
	if err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "gagal register user: "+err.Error())
	}

	return helper.JSONSuccessResponse(ctx, "Berhasil Register User")
}

func (c *AuthController) Login(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusBadRequest, "gagal mendapatkan data: "+err.Error())
	}

	token, err := c.AuthUsecase.Login(user.Email, user.Password)
	if err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusUnauthorized, "Login gagal: "+err.Error())
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   72 * 60 * 60,
	}

	ctx.SetCookie(cookie)

	return helper.JSONSuccessResponse(ctx, map[string]string{
		"Token": token,
	})
}

func (c *AuthController) Logout(ctx echo.Context) error {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   -1,
	}

	ctx.SetCookie(cookie)

	return helper.JSONSuccessResponse(ctx, "Berhasil Logout")
}
