package repository

import (
	"project_mini_golang/domain/model"

	"gorm.io/gorm"
)

type AirQualityRepository interface {
	GetAllData() ([]model.AirQualityData, error)
	GetDataByID(id int) (model.AirQualityData, error)
	CreateData(*model.AirQualityData) error
	UpdateData(id int, data *model.AirQualityData) error
	DeleteData(id int) error
}

type AirQualityRepositoryimpl struct {
	DB *gorm.DB
}

func NewAirQualityRepository(DB *gorm.DB) AirQualityRepository {
	return &AirQualityRepositoryimpl{DB: DB}
}

func (r *AirQualityRepositoryimpl) GetAllData() ([]model.AirQualityData, error) {
	var data []model.AirQualityData
	if err := r.DB.Preload("Alerts").Preload("Rekomendasi").Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *AirQualityRepositoryimpl) GetDataByID(id int) (model.AirQualityData, error) {
	var data model.AirQualityData
	if err := r.DB.Where("id =?", id).Preload("Alerts").Preload("Rekomendasi").First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (r *AirQualityRepositoryimpl) CreateData(data *model.AirQualityData) error {
	return r.DB.Create(&data).Error
}

func (r *AirQualityRepositoryimpl) UpdateData(id int, data *model.AirQualityData) error {
	if err := r.DB.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return r.DB.Preload("Alerts").Preload("Rekomendasi").First(data, id).Error
}

func (r *AirQualityRepositoryimpl) DeleteData(id int) error {
	var data model.AirQualityData
	return r.DB.Delete(&data, id).Error
}
