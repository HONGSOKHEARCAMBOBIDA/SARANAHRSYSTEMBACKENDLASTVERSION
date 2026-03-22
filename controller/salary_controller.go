package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateSalary(c *gin.Context) {
	id := c.Param("id")
	var updatesalary models.SalaryRequestUpdate

	// Bind JSON
	if err := c.ShouldBindJSON(&updatesalary); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	updatesalary.DailyRate = float64(updatesalary.BaseSalary) / float64(updatesalary.Workday)

	// Update record
	result := config.DB.Model(&models.Salary{}).Where("id = ?", id).Updates(updatesalary)
	if result.Error != nil {
		share.RespondError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	share.ResponeSuccess(c, 200, "កែប្រែបានជោគជ័យ")
}
