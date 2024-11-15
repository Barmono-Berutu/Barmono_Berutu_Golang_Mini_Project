package routes

import (
	"project_mini_golang/controllers"

	"github.com/labstack/echo/v4"
)

func AlertRoutes(e *echo.Group, airqualityroutes *controllers.AlertController) {
	e.GET("/alert", airqualityroutes.GetAllDataAlerts)
}
