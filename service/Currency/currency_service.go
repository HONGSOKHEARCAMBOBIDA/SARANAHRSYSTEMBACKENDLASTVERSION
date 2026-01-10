package service

import (
	"HRbackend/config"
	models "HRbackend/model"
	currency "HRbackend/request/Currency"
	"errors"

	"gorm.io/gorm"
)

type CurrencyService interface {
	Create(input currency.CurrencyRequestCreate) error
	GetAll() ([]models.Currency, error)
	Update(id int, input currency.CurrencyRequestUpdate) error
	ChangeStatus(id int) error
}

type currencyService struct {
	db *gorm.DB
}

func NewCurrencyService() CurrencyService {
	return &currencyService{
		db: config.DB,
	}
}

func (s *currencyService) Create(input currency.CurrencyRequestCreate) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	newCurrency := models.Currency{
		Name:     input.Name,
		Code:     input.Code,
		Symbol:   input.Symbol,
		Isactive: true,
	}
	if err := tx.Create(&newCurrency).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *currencyService) GetAll() ([]models.Currency, error) {
	var currency []models.Currency
	if err := s.db.Find(&currency).Error; err != nil {
		return nil, err
	}
	return currency, nil

}

func (s *currencyService) Update(id int, input currency.CurrencyRequestUpdate) error {
	result := s.db.Model(&models.Currency{}).Where("id =?", id).Updates(input)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("currency not found or no change")
	}
	return nil
}

func (s *currencyService) ChangeStatus(id int) error {
	result := s.db.Model(&models.Currency{}).Where("id =?", id).Update("is_active", gorm.Expr("!is_active"))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("currency not found or no change")
	}
	return nil
}
