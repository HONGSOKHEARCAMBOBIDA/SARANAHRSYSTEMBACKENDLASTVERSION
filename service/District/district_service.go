package service

import (
	"HRbackend/config"
	models "HRbackend/model"

	"gorm.io/gorm"
)

type DistrictService interface {
	GetByProvinceId(provinceid int) ([]models.District, error)
}
type districtService struct {
	db *gorm.DB
}

func NewDistrictService() DistrictService {
	return &districtService{
		db: config.DB,
	}
}
func (ds *districtService) GetByProvinceId(provinceid int) ([]models.District, error) {
	var district []models.District
	if err := ds.db.Where("province_id =?", provinceid).Find(&district).Error; err != nil {
		return nil, err
	}
	return district, nil
}
