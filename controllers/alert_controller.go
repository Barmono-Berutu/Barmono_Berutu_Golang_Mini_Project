package controllers

import (
	"net/http"
	"project_mini_golang/domain/usecase"
	"project_mini_golang/helper"

	"github.com/labstack/echo/v4"
)

type AlertController struct {
	AlertUsecase usecase.AlertUsecase
}

func NewAlertController(alertUsecase usecase.AlertUsecase) *AlertController {
	return &AlertController{
		AlertUsecase: alertUsecase,
	}
}

func (c *AlertController) GetAllDataAlerts(ctx echo.Context) error {
	alerts, err := c.AlertUsecase.GetAllAlerts()
	if err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "Gagal mengambil data alert: "+err.Error())
	}
	return helper.JSONSuccessResponse(ctx, alerts)
}
