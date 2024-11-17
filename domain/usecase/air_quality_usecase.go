package usecase

import (
	"errors"
	"project_mini_golang/domain/model"
	"project_mini_golang/domain/repository"
)

type AirQualityUsecase interface {
	GetData() ([]model.AirQualityData, error)
	GetDataID(id int) (model.AirQualityData, error)
	PostData(*model.AirQualityData) error
	PutData(id int, data *model.AirQualityData) error
	DeleteDatas(id int) error
}

type AirQualityUsecaseimpl struct {
	Repo repository.AirQualityRepository
}

func NewAirQualityUsecase(repo repository.AirQualityRepository) AirQualityUsecase {
	return &AirQualityUsecaseimpl{Repo: repo}
}

func (u *AirQualityUsecaseimpl) GetData() ([]model.AirQualityData, error) {
	return u.Repo.GetAllData()
}

func (u *AirQualityUsecaseimpl) GetDataID(id int) (model.AirQualityData, error) {
	return u.Repo.GetDataByID(id)
}

func (u *AirQualityUsecaseimpl) PostData(data *model.AirQualityData) error {
	return u.Repo.CreateData(data)
}
func (u *AirQualityUsecaseimpl) PutData(id int, data *model.AirQualityData) error {
	var dataasli model.AirQualityData
	if dataasli.UserID != data.UserID {
		return errors.New("user_id tidak boleh diubah")
	}
	return u.Repo.UpdateData(id, data)
}
func (u *AirQualityUsecaseimpl) DeleteDatas(id int) error {
	return u.Repo.DeleteData(id)
}
