package routes

import (
	"project_mini_golang/controllers"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group, authController *controllers.AuthController) {
	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)
}
