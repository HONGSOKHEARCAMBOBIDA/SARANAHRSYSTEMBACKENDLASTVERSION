package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Getuserv2(c *gin.Context) {
	var users []models.User

	branchID := c.Query("branch_id")
	roleID := c.Query("role_id")
	isActive := c.Query("is_active")

	db := config.DB.Preload("Branch").
		Preload("Role").
		Preload("Village.Communce.District.Province")

	if branchID != "" {
		db = db.Where("branch_id = ?", branchID)
	}
	if roleID != "" {
		db = db.Where("role_id = ?", roleID)
	}
	if isActive != "" {
		db = db.Where("is_active = ?", isActive)
	}

	// Pagination
	limit := 20
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	offset := (page - 1) * limit
	db = db.Limit(limit).Offset(offset)

	if err := db.Find(&users).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.RespondDate(c, http.StatusOK, users)
}
