package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	branch "HRbackend/request/Branch"
	branchresponse "HRbackend/response/Branch"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBranch(c *gin.Context) {
	var input branch.BranchRequestCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	tx := config.DB.Begin()
	if tx.Error != nil {
		share.RespondError(c, http.StatusInternalServerError, tx.Error.Error())
		return
	}

	// Ensure rollback on any panic or early return
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	newBranch := models.Branch{
		Name:      input.Name,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Radius:    input.Radius,
		IsActive:  1,
	}

	if err := tx.Create(&newBranch).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit().Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.ResponeSuccess(c, http.StatusOK, "Branch created")
}
func GetBranch(c *gin.Context) {
	var branch []branchresponse.BranchResponse
	if err := config.DB.Find(&branch).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, branch)
}
func UpdateBranch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		share.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	var updatebranch branch.BranchRequestUpdate
	if err := c.ShouldBindJSON(&updatebranch); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	result := config.DB.Model(&models.Branch{}).Where("id = ?", id).Updates(updatebranch)
	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusBadRequest, "សាខារកមិនឃេីញឬមិនមានការកែប្រែ")
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "សាខាត្រូវបានកែប្រែ")
}
func ChnageStatusBranch(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Model(&models.Branch{}).Where("id =?", id).
		Update("is_active", gorm.Expr("1 - is_active"))
	if result.Error != nil {
		share.RespondError(c, http.StatusNotFound, result.Error.Error())
		return
	}
	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusNotFound, "សាខារកមិនឃេីញឬមិនមានការកែប្រែ")
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "បានប្ដូស្ថានភាព")
}
