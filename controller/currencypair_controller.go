package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	currencypairgo "HRbackend/request/Currency_Pair.go"
	currencypairres "HRbackend/response/Currency_Pair.go"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCurrencyPair(c *gin.Context) {
	var input currencypairgo.CurrencyPairRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	tx := config.DB.Begin()
	if tx.Error != nil {
		share.RespondError(c, http.StatusInternalServerError, tx.Error.Error())
		return
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	newCurrencyPair := models.CurrencyPair{
		BaseCurrencyID:   input.BaseCurrencyID,
		TargetCurrencyID: input.TargetCurrencyID,
		IsActive:         true,
	}

	if err := tx.Create(&newCurrencyPair).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()
	share.ResponeSuccess(c, http.StatusOK, "Create Currency Success")
}

func GetCurrencypair(c *gin.Context) {
	var currencyPairs []currencypairres.CurrencyPairResponse

	db := config.DB.Table("currency_pairs").Select(`
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
		target.is_active AS target_currency_is_active ,
		currency_pairs.is_active AS is_active
	`).
		Joins("INNER JOIN currencies AS base ON base.id = currency_pairs.base_currency_id").
		Joins("INNER JOIN currencies AS target ON target.id = currency_pairs.target_currency_id").
		Order("currency_pairs.id DESC")

	if err := db.Scan(&currencyPairs).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.RespondDate(c, http.StatusOK, currencyPairs)
}
func UpdateCurrencyPaire(c *gin.Context) {
	id := c.Param("id")
	var updatecurrencypare currencypairgo.CurrencyPairRequest
	if err := c.ShouldBindJSON(&updatecurrencypare); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	result := config.DB.Model(&models.CurrencyPair{}).Where("id =?", id).Updates(updatecurrencypare)
	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusInternalServerError, "កែមិនបាន")
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "update success")
}

func ChangeStatusCurrencyPair(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Model(&models.CurrencyPair{}).Where("id =?", id).Update("is_active", gorm.Expr("!is_active"))
	if result.Error != nil {
		share.RespondError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "បានប្ដូស្ថានភាព")
}
