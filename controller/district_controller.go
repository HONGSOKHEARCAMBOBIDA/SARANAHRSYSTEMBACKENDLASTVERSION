package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDistrict(c *gin.Context) {
	proviceID := c.Param("id")
	var district []models.District
	if err := config.DB.Where("province_id =?", proviceID).Find(&district).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, district)
}
