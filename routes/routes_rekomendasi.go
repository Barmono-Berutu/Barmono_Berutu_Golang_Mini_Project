package routes

import (
	"project_mini_golang/controllers"

	"github.com/labstack/echo/v4"
)

func RekomendasiRoutes(e *echo.Group, airqualityroutes *controllers.RekomendasiController) {
	e.GET("/rekomendasi", airqualityroutes.GetAllDataAlerts)
}
