package model

import "time"

type Rekomendasi struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	AirQualityID int       `json:"air_quality_id" gorm:"constraint:OnDelete:CASCADE"`
	Message      string    `json:"message_rekomendasi"`
	CreatedAt    time.Time `json:"created_at"`
}
