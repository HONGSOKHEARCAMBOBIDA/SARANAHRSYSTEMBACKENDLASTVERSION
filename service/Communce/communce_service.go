package service

import (
	"HRbackend/config"
	models "HRbackend/model"

	"gorm.io/gorm"
)

type CommunceService interface {
	GetByDistrictId(districtID int) ([]models.Communce, error)
}

type communceService struct {
	db *gorm.DB
}

func NewCommunceService() CommunceService {
	return &communceService{
		db: config.DB,
	}
}

func (cm *communceService) GetByDistrictId(districtID int) ([]models.Communce, error) {
	var communces []models.Communce

	if err := cm.db.
		Where("district_id = ?", districtID).
		Find(&communces).Error; err != nil {
		return nil, err
	}

	return communces, nil
}
