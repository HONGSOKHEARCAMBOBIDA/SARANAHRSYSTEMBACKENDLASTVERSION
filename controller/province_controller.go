package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProvince(c *gin.Context) {
	var province []models.Province
	config.DB.Find(&province)
	share.RespondDate(c, http.StatusOK, province)
}
