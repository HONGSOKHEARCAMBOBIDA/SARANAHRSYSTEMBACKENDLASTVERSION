package currencypairser

import (
	"HRbackend/config"
	models "HRbackend/model"
	currencypairgo "HRbackend/request/Currency_Pair.go"
	currencypairres "HRbackend/response/Currency_Pair.go"
	"errors"

	"gorm.io/gorm"
)

type CurrencyPairService interface {
	CreateCurrencyPair(input currencypairgo.CurrencyPairRequest) error
	GetCurrencypair() ([]currencypairres.CurrencyPairResponse, error)
	UpdateCurrencyPaire(id int, input currencypairgo.CurrencyPairRequest) error
	ChangeStatusCurrencyPair(id int) error
}

type currencypairService struct {
	db *gorm.DB
}

func NewCurrencyPairService() CurrencyPairService {
	return &currencypairService{
		db: config.DB,
	}
}
func (cps *currencypairService) CreateCurrencyPair(input currencypairgo.CurrencyPairRequest) error {
	tx := cps.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	newcurrencypair := models.CurrencyPair{
		BaseCurrencyID:   input.BaseCurrencyID,
		TargetCurrencyID: input.TargetCurrencyID,
		IsActive:         true,
	}
	if err := tx.Create(&newcurrencypair).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (cps *currencypairService) GetCurrencypair() ([]currencypairres.CurrencyPairResponse, error) {
	var currencypair []currencypairres.CurrencyPairResponse

	db := cps.db.Table("currency_pairs").
		Select(`
			currency_pairs.id AS id,
			base.id AS base_currency_id,
			base.code AS base_currency_code,
			base.symbol AS base_currency_symbol,
			base.name AS base_currency_name,
			base.is_active AS base_currency_is_active,
			target.id AS target_currency_id,
			target.code AS target_currency_code,
			target.symbol AS target_currency_symbol,
			target.name AS target_currency_name,
			target.is_active AS target_currency_is_active,
			currency_pairs.is_active AS is_active
		`).
		Joins("INNER JOIN currencies AS base ON base.id = currency_pairs.base_currency_id").
		Joins("INNER JOIN currencies AS target ON target.id = currency_pairs.target_currency_id").
		Order("currency_pairs.id DESC")

	if err := db.Scan(&currencypair).Error; err != nil {
		return nil, err
	}

	return currencypair, nil
}

func (cps *currencypairService) UpdateCurrencyPaire(id int, input currencypairgo.CurrencyPairRequest) error {
	result := cps.db.Model(&models.CurrencyPair{}).Where("id =?", id).Updates(input)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("currency pair not foundn or not update")
	}
	return nil
}

func (cps *currencypairService) ChangeStatusCurrencyPair(id int) error {
	result := cps.db.Model(&models.CurrencyPair{}).Where("id =?", id).Update("is_active", gorm.Expr("!is_active"))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("currency pair not found or not update")

	}
	return nil
}
