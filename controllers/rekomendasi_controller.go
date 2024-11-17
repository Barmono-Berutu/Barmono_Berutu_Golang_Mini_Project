package controllers

import (
	"net/http"
	"project_mini_golang/domain/usecase"
	"project_mini_golang/helper"

	"github.com/labstack/echo/v4"
)

type RekomendasiController struct {
	RekomendasiUsecase usecase.RekomendasiUsecase
}

func NewRekomendasiController(rekomendasiUsecase usecase.RekomendasiUsecase) *RekomendasiController {
	return &RekomendasiController{
		RekomendasiUsecase: rekomendasiUsecase,
	}
}

func (c *RekomendasiController) GetAllDataAlerts(ctx echo.Context) error {
	alerts, err := c.RekomendasiUsecase.GetAllRekomendasi()
	if err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "Gagal mengambil data alert: "+err.Error())
	}
	return helper.JSONSuccessResponse(ctx, alerts)
}
