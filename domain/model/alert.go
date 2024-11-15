package model

import "time"

type Alert struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	AirQualityID int       `json:"air_quality_id" gorm:"constraint:OnDelete:CASCADE"`
	AlertMessage string    `json:"alert_message"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}
