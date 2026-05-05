package service

import (
	"HRbackend/config"
	models "HRbackend/model"
	branch "HRbackend/request/Branch"
	"errors"

	"gorm.io/gorm"
)

type BranchService interface {
	Create(intput branch.BranchRequestCreate) error
	GetAll() ([]models.Branch, error)
	Update(id int, input branch.BranchRequestUpdate) error
	ChangeStatus(id int) error
}
type branchService struct {
	db *gorm.DB
}

func NewBranchService() BranchService {
	return &branchService{
		db: config.DB,
	}
}
func (s *branchService) Create(input branch.BranchRequestCreate) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	newbranch := models.Branch{
		Name:      input.Name,
		Latitude:  *input.Latitude,
		Longitude: *input.Longitude,
		Radius:    input.Radius,
		IsActive:  1,
	}
	if err := tx.Create(&newbranch).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (s *branchService) GetAll() ([]models.Branch, error) {
	var branches []models.Branch
	if err := s.db.Order("id DESC").Find(&branches).Error; err != nil {
		return nil, err
	}
	return branches, nil
}
func (s *branchService) Update(id int, input branch.BranchRequestUpdate) error {
	result := s.db.Model(&models.Branch{}).
		Where("id = ?", id).
		Updates(input)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("branch not found or no changes")
	}

	return nil
}
func (s *branchService) ChangeStatus(id int) error {
	result := s.db.Model(&models.Branch{}).
		Where("id = ?", id).
		Update("is_active", gorm.Expr("1 - is_active"))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("branch not found")
	}

	return nil
}
