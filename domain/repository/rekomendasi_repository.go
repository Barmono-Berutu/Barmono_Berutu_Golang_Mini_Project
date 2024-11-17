package repository

import (
	"project_mini_golang/domain/model"

	"gorm.io/gorm"
)

type RekomendasiRepository interface {
	GetRekomendasi() ([]model.Rekomendasi, error)
	CreateRekomendasi(rekomendasi *model.Rekomendasi) error
	FindRekomendasiByAirQualityID(airQualityID int) (*model.Rekomendasi, error)
	UpdateRekomendasi(rekomendasi *model.Rekomendasi) error
}

type RekomendasiRepositoryimpl struct {
	DB *gorm.DB
}

func NewRekomendasiRepository(db *gorm.DB) RekomendasiRepository {
	return &RekomendasiRepositoryimpl{DB: db}
}

func (r *RekomendasiRepositoryimpl) GetRekomendasi() ([]model.Rekomendasi, error) {
	var rekomendasi []model.Rekomendasi
	return rekomendasi, r.DB.Find(&rekomendasi).Error
}

func (r *RekomendasiRepositoryimpl) CreateRekomendasi(rekomendasi *model.Rekomendasi) error {
	return r.DB.Create(rekomendasi).Error
}

func (r *RekomendasiRepositoryimpl) FindRekomendasiByAirQualityID(airQualityID int) (*model.Rekomendasi, error) {
	var rekomendasi model.Rekomendasi
	err := r.DB.Where("air_quality_id = ?", airQualityID).First(&rekomendasi).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &rekomendasi, err
}

func (r *RekomendasiRepositoryimpl) UpdateRekomendasi(rekomendasi *model.Rekomendasi) error {
	return r.DB.Save(rekomendasi).Error
}
