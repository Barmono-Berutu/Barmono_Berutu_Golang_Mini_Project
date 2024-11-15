package controllers

import (
	"net/http"
	"project_mini_golang/domain/model"
	"project_mini_golang/domain/usecase"
	"project_mini_golang/helper"

	"github.com/labstack/echo/v4"
)

type AirQualityController struct {
	AirQualityUsecase  usecase.AirQualityUsecase
	AlertUsecase       usecase.AlertUsecase
	RekomendasiUsecase usecase.RekomendasiUsecase
}

func NewAirQualityController(air usecase.AirQualityUsecase, alertusecase usecase.AlertUsecase, rekomendasiusecase usecase.RekomendasiUsecase) *AirQualityController {
	return &AirQualityController{
		AirQualityUsecase:  air,
		AlertUsecase:       alertusecase,
		RekomendasiUsecase: rekomendasiusecase,
	}
}

func (c *AirQualityController) GetData(ctx echo.Context) error {
	data, err := c.AirQualityUsecase.GetData()
	if err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "Gagal mengambil data kualitas udara: "+err.Error())
	}
	return helper.JSONSuccessResponse(ctx, data)
}

func (c *AirQualityController) GetDataByID(ctx echo.Context) error {
	id, err := helper.GetIDParam(ctx)
	if err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusBadRequest, "Format ID tidak valid: "+err.Error())
	}

	data, err := c.AirQualityUsecase.GetDataID(id)
	if err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "Gagal mengambil data kualitas udara: "+err.Error())
	}
	return helper.JSONSuccessResponse(ctx, data)
}

func (c *AirQualityController) PostData(ctx echo.Context) error {
	userID := ctx.Get("user_id").(int)
	data := new(model.AirQualityData)
	if err := ctx.Bind(data); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusBadRequest, "Gagal menginput data: "+err.Error())
	}

	data.UserID = userID

	if err := c.AirQualityUsecase.PostData(data); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "Gagal menyimpan data kualitas udara: "+err.Error())
	}

	if err := c.AlertUsecase.GenerateAlert(ctx.Request().Context(), data); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "Gagal menghasilkan alert: "+err.Error())
	}

	if err := c.RekomendasiUsecase.GenerateRekomendasi(ctx.Request().Context(), data); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "Gagal menghasilkan rekomendasi: "+err.Error())
	}

	return helper.JSONSuccessResponse(ctx, "Data kualitas udara berhasil ditambahkan")
}

func (c *AirQualityController) UpdateData(ctx echo.Context) error {
	id, err := helper.GetIDParam(ctx)
	if err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusBadRequest, "Format ID tidak valid: "+err.Error())
	}

	data := new(model.AirQualityData)
	if err := ctx.Bind(data); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusBadRequest, "Gagal mengikat data: "+err.Error())
	}

	if err := c.AirQualityUsecase.PutData(id, data); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "Gagal memperbarui data udara: "+err.Error())
	}

	if err := c.AlertUsecase.GenerateAlert(ctx.Request().Context(), data); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "Gagal memperbarui alert: "+err.Error())
	}

	if err := c.RekomendasiUsecase.GenerateRekomendasi(ctx.Request().Context(), data); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "Gagal memperbarui rekomendasi: "+err.Error())
	}

	return helper.JSONSuccessResponse(ctx, map[string]interface{}{
		"message": "Data udara berhasil diperbarui dan respons AI telah diperbarui",
		"data":    data,
	})

}

func (c *AirQualityController) DeleteData(ctx echo.Context) error {
	id, err := helper.GetIDParam(ctx)
	if err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusBadRequest, "Format ID tidak valid: "+err.Error())
	}

	if err := c.AirQualityUsecase.DeleteDatas(id); err != nil {
		return helper.JSONErrorResponse(ctx, http.StatusInternalServerError, "Gagal menghapus data udara: "+err.Error())
	}

	return helper.JSONSuccessResponse(ctx, "Data udara berhasil dihapus")
}
