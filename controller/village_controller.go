package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVillage(c *gin.Context) {
	CommunceId := c.Param("id")
	var village []models.Village
	if err := config.DB.Where("communce_id =?", CommunceId).Find(&village).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, village)
}
