package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateShift(c *gin.Context) {
	var input models.ShiftRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	tx := config.DB.Begin()
	if tx.Error != nil {
		share.RespondError(c, http.StatusInternalServerError, tx.Error.Error())
		return
	}

	newshift := models.Shift{
		Name:      input.Name,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		BranchID:  input.BranchID,
		IsActive:  1,
	}

	if err := tx.Create(&newshift).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()
	share.ResponeSuccess(c, http.StatusOK, "shift created")
}

func GetShift(c *gin.Context) {
	var shifts []models.Shift
	if err := config.DB.Find(&shifts).Error; err != nil {
		share.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	share.RespondDate(c, http.StatusOK, shifts)
}

func GetShiftByBranchID(c *gin.Context) {
	branchID := c.Param("id")
	var shifts []models.Shift
	if err := config.DB.Where("branch_id =?", branchID).Find(&shifts).Error; err != nil {
		share.RespondError(c, http.StatusNotFound, err.Error())
		return
	}
	// var resp []models.ShiftResponse
	// for _, s := range shifts {
	// 	resp = append(resp, models.ShiftResponse{
	// 		ID:        s.ID,
	// 		Name:      s.Name + " " + s.StartTime + " - " + s.EndTime,
	// 		StartTime: s.StartTime,
	// 		EndTime:   s.EndTime,
	// 		IsActive:  s.IsActive,
	// 		BranchID:  s.BranchID,
	// 	})
	// }
	share.RespondDate(c, http.StatusOK, shifts)
}

func UpdateShift(c *gin.Context) {
	id := c.Param("id")
	var updateshift models.ShiftRequest
	if err := c.ShouldBindJSON(&updateshift); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	result := config.DB.Model(&models.Shift{}).Where("id =?", id).Updates(&updateshift)
	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusBadRequest, "ម៉ោងការងាររកមិនឃេីញឬមិនបានការកែប្រេ")
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "ម៉ោងការងារកែប្រែបានសម្រេច")
}

func ChangeStatusShift(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Model(&models.Shift{}).Where("id =?", id).
		Update("is_active", gorm.Expr("1 - is_active"))
	if result.Error != nil {
		share.RespondError(c, http.StatusNotFound, result.Error.Error())
		return
	}
	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusNotFound, "ម៉ោងការងាររកមិនឃេីញ")
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "បានប្ដូស្ថានភាព")
}
