package usecase

import (
	"context"
	"fmt"
	"project_mini_golang/domain/model"
	"project_mini_golang/domain/repository"
	"project_mini_golang/helper"
	"time"
)

type AlertUsecase interface {
	GenerateAlert(ctx context.Context, data *model.AirQualityData) error
	GetAllAlerts() ([]model.Alert, error)
}

type AlertUsecaseImpl struct {
	Repo repository.AlertRepository
}

func NewAlertUsecase(repo repository.AlertRepository) AlertUsecase {
	return &AlertUsecaseImpl{Repo: repo}
}

func (u *AlertUsecaseImpl) GenerateAlert(ctx context.Context, data *model.AirQualityData) error {
	pertanyaan := fmt.Sprintf(`Data kualitas udara:
	- Lokasi: %s
	- PM2.5: %.2f
	- CO: %.2f
	- O3: %.2f
	- Risk Index: %.2f
	
	Berikan saran yang sesuai berdasarkan data tersebut dan pastikan untuk menyebutkan alert level dan pesan peringatan.`,
		data.Location, data.PM25, data.CO, data.O3, data.RiskIndex)

	jawaban, err := helper.ResponseAI(ctx, pertanyaan)
	if err != nil {
		return err
	}

	existingAlert, err := u.Repo.FindAlertByAirQualityID(data.ID)
	if err != nil {
		return err
	}

	if existingAlert != nil {
		existingAlert.AlertMessage = jawaban
		existingAlert.ExpiresAt = time.Now().Add(24 * time.Hour)
		return u.Repo.UpdateAlert(existingAlert)
	}

	newAlert := &model.Alert{
		AirQualityID: data.ID,
		AlertMessage: jawaban,
		CreatedAt:    time.Now(),
		ExpiresAt:    time.Now().Add(24 * time.Hour),
	}

	return u.Repo.CreateAlert(newAlert)
}

func (u *AlertUsecaseImpl) GetAllAlerts() ([]model.Alert, error) {
	return u.Repo.GetAlerts()
}
