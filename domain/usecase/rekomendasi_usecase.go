package usecase

import (
	"context"
	"fmt"
	"project_mini_golang/domain/model"
	"project_mini_golang/domain/repository"
	"project_mini_golang/helper"
	"time"
)

type RekomendasiUsecase interface {
	GenerateRekomendasi(ctx context.Context, data *model.AirQualityData) error
	GetAllRekomendasi() ([]model.Rekomendasi, error)
}

type RekomendasiUsecaseImpl struct {
	Repo repository.RekomendasiRepository
}

func NewRekomendasiUsecase(repo repository.RekomendasiRepository) RekomendasiUsecase {
	return &RekomendasiUsecaseImpl{Repo: repo}
}

func (u *RekomendasiUsecaseImpl) GenerateRekomendasi(ctx context.Context, data *model.AirQualityData) error {

	pertanyaan := fmt.Sprintf("Berikan rekomendasi kesehatan dari data udara berikut: %+v", data)

	jawaban, err := helper.ResponseAI(ctx, pertanyaan)
	if err != nil {
		return err
	}

	// Cek apakah sudah ada alert untuk AirQualityID ini
	existingRekomendasi, err := u.Repo.FindRekomendasiByAirQualityID(data.ID)
	if err != nil {
		return err
	}

	if existingRekomendasi != nil {
		// Update alert jika sudah ada
		existingRekomendasi.Message = jawaban
		existingRekomendasi.CreatedAt = time.Now().Add(24 * time.Hour)
		return u.Repo.UpdateRekomendasi(existingRekomendasi)
	}

	// Jika belum ada alert, buat alert baru
	newRekomendasi := &model.Rekomendasi{
		AirQualityID: data.ID,
		Message:      jawaban,
		CreatedAt:    time.Now(),
	}

	return u.Repo.CreateRekomendasi(newRekomendasi)
}

func (u *RekomendasiUsecaseImpl) GetAllRekomendasi() ([]model.Rekomendasi, error) {
	return u.Repo.GetRekomendasi()
}
