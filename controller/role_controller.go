package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRole(c *gin.Context) {
	var role []models.Role
	if err := config.DB.Find(&role).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.RespondDate(c, http.StatusOK, role)
}
func CreateRole(c *gin.Context) {
	var input models.RoleRequestCreate

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Map input to Role model
	newRole := models.Role{
		Name:        input.Name,
		DisPlayName: input.DisPlayName,
		IsActive:    true, // Respect input value
	}

	// Create role
	if err := config.DB.Create(&newRole).Error; err != nil {
		share.RespondError(c, 500, err.Error())
		return
	}

	// Success response
	share.ResponeSuccess(c, 200, "បង្កើតតួនាទីបានជោគជ័យ")
}
func ChangeStatusRole(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	result := config.DB.Model(&models.Role{}).Where("id =?", id).Update("is_active", gorm.Expr("1 - is_active"))
	if result.Error != nil {
		share.RespondError(c, 500, result.Error.Error())
		return
	}
	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusNotFound, result.Error.Error())
		return
	}
	share.ResponeSuccess(c, 200, "កែប្រែស្ថានភាពបានជោគជ័យ")
}
func UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var updaterole models.RoleRequestUpdate
	if err := c.ShouldBindJSON(&updaterole); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	result := config.DB.Model(&models.Role{}).Where("id =?", id).Updates(updaterole)
	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusBadRequest, "អ្នកប្រេីប្រាស់រកមិនឃេីញឬមិនមានអ្វីត្រូវកែ")
		return
	}

	share.ResponeSuccess(c, http.StatusOK, "តួនាទីប្រេីប្រាស់ត្រូវបានកែប្រែ")
}
