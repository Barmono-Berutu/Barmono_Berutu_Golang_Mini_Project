package model

// Model untuk User
type User struct {
	ID             int              `json:"id" gorm:"primaryKey"`
	Email          string           `json:"email"`
	Password       string           `json:"password"`
	AirQualityData []AirQualityData `gorm:"constraint:OnDelete:CASCADE;"`
}
