package routes

import (
	"project_mini_golang/controllers"

	"github.com/labstack/echo/v4"
)

func AirqualityRoutes(e *echo.Group, airqualityroutes *controllers.AirQualityController) {
	e.GET("", airqualityroutes.GetData)
	e.GET("/:id", airqualityroutes.GetDataByID)
	e.POST("", airqualityroutes.PostData)
	e.PUT("/:id", airqualityroutes.UpdateData)
	e.DELETE("/:id", airqualityroutes.DeleteData)
}
