package service

import (
	"HRbackend/config"
	models "HRbackend/model"

	"gorm.io/gorm"
)

type VillageService interface {
	GetByCommunceId(communceid int) ([]models.Village, error)
}
type villageService struct {
	db *gorm.DB
}

func NewVillageService() VillageService {
	return &villageService{
		db: config.DB,
	}
}
func (vs *villageService) GetByCommunceId(communceid int) ([]models.Village, error) {
	var villagees []models.Village
	if err := vs.db.Where("communce_id =?", communceid).Find(&villagees).Error; err != nil {
		return nil, err
	}
	return villagees, nil
}
