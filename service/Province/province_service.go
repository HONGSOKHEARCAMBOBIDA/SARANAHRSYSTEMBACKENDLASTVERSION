package province

import (
	"HRbackend/config"
	models "HRbackend/model"

	"gorm.io/gorm"
)

type ProvinceService interface {
	GetAll() ([]models.Province, error)
}
type provinceService struct {
	db *gorm.DB
}

func NewProvinceService() ProvinceService {
	return &provinceService{
		db: config.DB,
	}
}
func (pc *provinceService) GetAll() ([]models.Province, error) {
	var province []models.Province
	if err := pc.db.Find(&province).Error; err != nil {
		return nil, err
	}
	return province, nil
}
