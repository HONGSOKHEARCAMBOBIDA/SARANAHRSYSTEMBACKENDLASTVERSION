package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	currency "HRbackend/request/Currency"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCurrency(c *gin.Context) {
	var input currency.CurrencyRequestCreate
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
	newCurrency := models.Currency{
		Code:     input.Code,
		Symbol:   input.Symbol,
		Name:     input.Name,
		Isactive: true,
	}
	if err := tx.Create(&newCurrency).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	tx.Commit()
	share.ResponeSuccess(c, http.StatusOK, "Currency Created")
}
func GetCurrency(c *gin.Context) {
	var currency []models.Currency
	if err := config.DB.Find(&currency).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, currency)
}
func UpdateCurrency(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	var updatecurrency currency.CurrencyRequestUpdate
	if err := c.ShouldBindJSON(&updatecurrency); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	result := config.DB.Model(&models.Currency{}).Where("id =?", id).Updates(updatecurrency)
	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusInternalServerError, "កែមិនបាន")
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "udpate currrency success")
}
func ChangeStatusCurrency(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Model(&models.Currency{}).Where("id =?", id).Update("is_active", gorm.Expr("!is_active"))
	if result.Error != nil {
		share.RespondError(c, http.StatusNotFound, result.Error.Error())
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "បានប្ដូស្ថានភាព")
}
