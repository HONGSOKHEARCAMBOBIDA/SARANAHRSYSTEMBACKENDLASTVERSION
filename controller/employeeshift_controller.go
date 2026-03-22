package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	"HRbackend/helper"
	models "HRbackend/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateEmployeeShift(c *gin.Context) {
	id := c.Param("id")
	var updateemployeeshift models.EmployeeShiftRequestUpdate
	if err := c.ShouldBindJSON(&updateemployeeshift); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	result := config.DB.Model(&models.EmployeeShift{}).Where("id =?", id).Updates(updateemployeeshift)
	if result.Error != nil {
		share.RespondError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	share.ResponeSuccess(c, 200, "កែប្រែបានជោគជ័យ")
}
func CreateEmployeeShift(c *gin.Context) {
	employeeShiftID := c.Param("employeeshiftid")
	salaryID := c.Param("salaryid")

	var input struct {
		EmployeeID     int       `json:"employee_id" binding:"required"`
		ShiftID        int       `json:"shift_id" binding:"required"`
		AssignBranchID int       `json:"assign_branch_id"`
		BaseSalary     float64   `json:"base_salary" binding:"required"`
		WorkedDay      int       `json:"worked_day" binding:"required"`
		EffectiveDate  time.Time `json:"effective_date"`
		ExpireDate     time.Time `json:"expire_date" `
		CurrencyID     int       `json:"currency_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	tx := config.DB.Begin()

	// 1️⃣ Deactivate old employee_shift
	if err := tx.Model(&models.EmployeeShift{}).
		Where("id = ?", employeeShiftID).
		Update("is_active", 0).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, "បរាជ័យក្នុងការកែស្ថានភាព shift ចាស់: "+err.Error())
		return
	}

	// 2️⃣ Set old salary inactive and update effective_date
	if err := tx.Model(&models.Salary{}).
		Where("id = ?", salaryID).
		Updates(map[string]interface{}{
			"is_active":   0,
			"expire_date": time.Now(),
		}).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, "បរាជ័យក្នុងការកែប្រាក់ខែចាស់: "+err.Error())
		return
	}

	// 3️⃣ Create new employee_shift
	newShift := models.EmployeeShift{
		EmployeeID:     input.EmployeeID,
		ShiftID:        input.ShiftID,
		AssignBranchID: input.AssignBranchID,
	}

	if err := tx.Create(&newShift).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, "បរាជ័យក្នុងការបង្កើត shift ថ្មី: "+err.Error())
		return
	}

	// 4️⃣ Create new salary
	newSalary := models.Salary{
		EmployeeShiftID: newShift.ID,
		BaseSalary:      input.BaseSalary,
		WorkedDay:       input.WorkedDay,
		DailyRate:       float64(input.BaseSalary) / float64(input.WorkedDay),
		EffectTiveDate:  time.Now(),
		CurrencyID:      input.CurrencyID,
	}

	if err := tx.Create(&newSalary).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, "បរាជ័យក្នុងការបង្កើតប្រាក់ខែថ្មី: "+err.Error())
		return
	}

	tx.Commit()
	share.ResponeSuccess(c, http.StatusOK, "បានកែប្រែម៉ោងធ្វើការ និងប្រាក់ខែដោយជោគជ័យ")
}
func GetEmployeeShiftByuserLogin(c *gin.Context) {
	var employeeshift []models.EmployeeShiftResponse

	userID, ok := helper.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please login!"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		share.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	db := config.DB.Table("employee_shifts").Select(`
		employee_shifts.id AS id,
		shifts.id AS shift_id,
		shifts.name AS shift_name,
		shifts.start_time AS start_time,
		shifts.end_time AS end_time
	`).
		Joins("INNER JOIN shifts ON shifts.id = employee_shifts.shift_id").
		Where("employee_shifts.employee_id = ?", user.EmployeeID).
		Where("employee_shifts.is_active = ?", 1)

	if err := db.Scan(&employeeshift).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.RespondDate(c, http.StatusOK, employeeshift)
}
