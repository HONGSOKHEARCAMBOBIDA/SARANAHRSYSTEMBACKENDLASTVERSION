package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCommunes(c *gin.Context) {

	districtIDStr := c.Param("id")

	// Convert to int (optional but safe)
	districtID, err := strconv.Atoi(districtIDStr)
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "Invalid district ID")
		return
	}

	var communes []models.Communce

	if err := config.DB.Where("district_id = ?", districtID).Find(&communes).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Optional: check empty
	if len(communes) == 0 {
		share.RespondDate(c, http.StatusOK, []models.Communce{})
		return
	}

	share.RespondDate(c, http.StatusOK, communes)
}
