package repository

import (
	"project_mini_golang/domain/model"

	"gorm.io/gorm"
)

type AlertRepository interface {
	GetAlerts() ([]model.Alert, error)
	CreateAlert(alert *model.Alert) error
	FindAlertByAirQualityID(airQualityID int) (*model.Alert, error)
	UpdateAlert(alert *model.Alert) error
}

type AlertRepositoryimpl struct {
	DB *gorm.DB
}

func NewAlertRepository(db *gorm.DB) AlertRepository {
	return &AlertRepositoryimpl{DB: db}
}

func (r *AlertRepositoryimpl) GetAlerts() ([]model.Alert, error) {
	var alerts []model.Alert
	return alerts, r.DB.Find(&alerts).Error
}

func (r *AlertRepositoryimpl) CreateAlert(alert *model.Alert) error {
	return r.DB.Create(alert).Error
}

// FindAlertByAirQualityID mencari alert berdasarkan AirQualityID
func (r *AlertRepositoryimpl) FindAlertByAirQualityID(airQualityID int) (*model.Alert, error) {
	var alert model.Alert
	err := r.DB.Where("air_quality_id = ?", airQualityID).First(&alert).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Tidak ada data
	}
	return &alert, err
}

// UpdateAlert memperbarui alert yang ada
func (r *AlertRepositoryimpl) UpdateAlert(alert *model.Alert) error {
	return r.DB.Save(alert).Error
}
