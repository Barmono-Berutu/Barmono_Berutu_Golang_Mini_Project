package model

import "time"

// Model untuk AirQualityData
type AirQualityData struct {
	ID          int           `json:"id" gorm:"primaryKey"`
	UserID      int           `json:"user_id" gorm:"constraint:OnDelete:CASCADE"`
	Location    string        `json:"location"`
	PM25        float64       `json:"pm2_5"`
	CO          float64       `json:"co"`
	O3          float64       `json:"o3"`
	RecordedAt  time.Time     `json:"recorded_at"`
	RiskIndex   float64       `json:"risk_index"`
	Status      string        `json:"status"`
	Alerts      []Alert       `json:"alerts" gorm:"foreignKey:AirQualityID;constraint:OnDelete:CASCADE"`
	Rekomendasi []Rekomendasi `json:"rekomendasi" gorm:"foreignKey:AirQualityID;constraint:OnDelete:CASCADE"`
}
