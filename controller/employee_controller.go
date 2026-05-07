package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	"HRbackend/helper"
	models "HRbackend/model"
	employeerequstupdate "HRbackend/request/Employee"
	employeeres "HRbackend/response/Employee"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetEmployee(c *gin.Context) {
	var employees []employeeres.EmployeeResponse

	branchID := c.Query("branch_id")
	name := c.Query("name")
	roleID := c.Query("role_id")
	isActive := c.Query("is_active")
	shiftID := c.Query("shift_id")
	isPromote := c.Query("is_promote")

	db := config.DB.Table("employees").Select(`
        employees.id AS id,
        branches.id AS branch_id,
        branches.name AS branch_name,
        employees.name_en AS name_en,
        employees.name_kh AS name_kh,
        employees.gender AS gender,
        employees.contact AS contact,
        employees.national_id_number AS national_id_number,
        employees.is_active AS is_active,
		employee_profiles.date_of_birth AS date_of_birth,
		employee_profiles.marital_status AS marital_status,
		employee_profiles.profile_image,
		employee_profiles.village_id_current_address,
		employee_profiles.family_phone AS family_phone,
		employee_profiles.education_level AS education_level,
		employee_profiles.experience_years AS experience_years,
		employee_profiles.previous_company AS previous_company,
		employee_profiles.bank_name AS bank_name,
		employee_profiles.bank_account_number AS bank_account_number,
		employee_profiles.qr_code_bank_account AS qr_code_bank_account,
		employee_profiles.notes AS notes,
		employee_profiles.position_level AS position_level,
		employee_shifts.assign_branch_id AS assign_branch_id,
		birth_village.id AS village_id_of_birth,
		birth_village.name AS village_name_of_birth,
		birth_communce.id AS communce_id_of_birth,
		birth_communce.name AS communce_name_of_birth,
		birth_district.id AS district_id_of_birth,
		birth_district.name AS district_name_of_birth,
        birth_province.id AS province_id_of_birth,
		birth_province.name AS province_name_of_birth,
		current_province.id AS province_id_current_address,
		current_province.name AS province_name_current_address,
		current_district.id AS district_id_current_address,
		current_district.name AS district_name_current_address,
		current_communce.id AS communce_id_current_address,
		current_communce.name AS communce_name_current_address,
		current_village.id AS village_id_current_address,
		current_village.name AS village_name_current_address,
        roles.id AS role_id,
        roles.display_name AS role_name,
        employees.hire_date AS hire_date,
		employees.promote_date AS promote_date,
		employees.is_promote AS is_promote,
        employees.type AS type,
        shifts.id AS shift_id,
        shifts.name AS shift_name,
        shifts.start_time AS start_time,
        shifts.end_time AS end_time,
		shifts.branch_id AS branch_shift_id,
        employee_shifts.id AS employee_shift_id,
        salaries.id AS salary_id,
        salaries.base_salary AS base_salary,
        salaries.worked_day AS worked_day,
        salaries.daily_rate AS daily_rate,
		currencies.id AS currency_id,
		currencies.code AS currency_code,
		currencies.symbol AS currency_symbol,
		currencies.name AS currency_name

    `).
		Joins("LEFT JOIN employee_profiles ON employee_profiles.employee_id = employees.id").
		Joins("LEFT JOIN villages AS birth_village ON birth_village.id = employee_profiles.village_id_of_birth").
		Joins("LEFT JOIN communces AS birth_communce ON birth_communce.id = birth_village.communce_id").
		Joins("LEFT JOIN districts AS birth_district ON birth_district.id = birth_communce.district_id").
		Joins("LEFT JOIN provinces AS birth_province ON birth_province.id = birth_district.province_id").
		Joins("LEFT JOIN villages AS current_village ON current_village.id = employee_profiles.village_id_current_address").
		Joins("LEFT JOIN communces AS current_communce ON current_communce.id = current_village.communce_id").
		Joins("LEFT JOIN districts AS current_district ON current_district.id = current_communce.district_id").
		Joins("LEFT JOIN provinces AS current_province ON current_province.id = current_district.province_id").
		Joins("LEFT JOIN branches ON branches.id = employees.branch_id").
		Joins("LEFT JOIN roles ON roles.id = employees.role_id").
		Joins("LEFT JOIN employee_shifts ON employee_shifts.employee_id = employees.id AND employee_shifts.is_active = 1").
		Joins("LEFT JOIN shifts ON shifts.id = employee_shifts.shift_id").
		Joins("LEFT JOIN salaries ON salaries.employee_shift_id = employee_shifts.id AND salaries.is_active = 1").
		Joins("LEFT JOIN currencies ON currencies.id = salaries.currency_id")

	if branchID != "" {
		db = db.Where("employees.branch_id = ?", branchID)
	}
	if name != "" {
		db = db.Where("employees.name_en LIKE ? OR employees.name_kh LIKE ?", "%"+name+"%", "%"+name+"%")
	}
	if roleID != "" {
		db = db.Where("employees.role_id = ?", roleID)
	}
	if isActive != "" {
		db = db.Where("employees.is_active = ?", isActive)
	}
	if shiftID != "" {
		db = db.Where("employee_shifts.shift_id = ?", shiftID)
	}
	if isPromote != "" {
		db = db.Where("employees.is_promote =?", isPromote)
	}
	db = db.Order("employees.id DESC")

	if err := db.Scan(&employees).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	for i := range employees {
		employees[i].DateOfBirth = helper.FormatDate(employees[i].DateOfBirth)
		employees[i].HireDate = helper.FormatDate(employees[i].HireDate)
		employees[i].PromoteDate = helper.FormatDate(employees[i].PromoteDate)
	}

	share.RespondDate(c, http.StatusOK, employees)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var updateemployee employeerequstupdate.EmployeeRequestUpdate

	var employeeprofile models.EmployeeProfile

	// Parse form data
	if err := c.ShouldBind(&updateemployee); err != nil {

		share.RespondError(c, http.StatusBadRequest, err.Error())

		return
	}

	// Find employee profile
	if err := config.DB.Where("employee_id = ?", id).First(&employeeprofile).Error; err != nil {

		share.RespondError(c, http.StatusNotFound, "Employee profile not found")

		return
	}

	// Update Employee main info
	result := config.DB.Model(&models.Employee{}).Where("id = ?", id).Updates(map[string]interface{}{

		"branch_id": updateemployee.BranchID,

		"name_en": updateemployee.NameEn,

		"name_kh": updateemployee.NameKh,

		"gender": updateemployee.Gender,

		"contact": updateemployee.Contact,

		"national_id_number": updateemployee.NationalIDNumber,

		"role_id": updateemployee.RoleID,

		"hire_date": updateemployee.HireDate,

		"promote_date": updateemployee.PromoteDate,

		"type": updateemployee.Type,
	})

	if result.Error != nil {

		share.RespondError(c, http.StatusInternalServerError, "Failed to update employee")

		return
	}

	// Handle profile image upload

	file, err := c.FormFile("profileimage")

	if err == nil {

		if !helper.ProtectImage(file) {

			share.RespondError(c, http.StatusBadRequest, "Image Not Allow")

			return
		}

		oldProfilePath := filepath.Join("public/profileimage", employeeprofile.ProfileImage)

		if _, err := os.Stat(oldProfilePath); err == nil {

			os.Remove(oldProfilePath)

		}

		profileImageDir := "public/profileimage"

		if _, err := os.Stat(profileImageDir); os.IsNotExist(err) {

			os.MkdirAll(profileImageDir, os.ModePerm)

		}

		extension := filepath.Ext(file.Filename)

		newProfileImageName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)

		profileImagePath := filepath.Join(profileImageDir, newProfileImageName)

		if err := c.SaveUploadedFile(file, profileImagePath); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload new image"})

			return
		}

		employeeprofile.ProfileImage = newProfileImageName
	}

	// Handle QR code image upload
	qrFile, err := c.FormFile("qrcodeimage")

	if err == nil {

		if !helper.ProtectImage(qrFile) {

			share.RespondError(c, http.StatusBadRequest, "Image Not Allow")

			return
		}

		oldQRPath := filepath.Join("public/qrcodeimage", employeeprofile.QrCodeBankAccount)

		if _, err := os.Stat(oldQRPath); err == nil {

			os.Remove(oldQRPath)

		}

		qrImageDir := "public/qrcodeimage"

		if _, err := os.Stat(qrImageDir); os.IsNotExist(err) {

			os.MkdirAll(qrImageDir, os.ModePerm)

		}

		extension := filepath.Ext(qrFile.Filename)

		newQRImageName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)

		qrImagePath := filepath.Join(qrImageDir, newQRImageName)

		if err := c.SaveUploadedFile(qrFile, qrImagePath); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload new QR image"})

			return

		}

		employeeprofile.QrCodeBankAccount = newQRImageName
	}

	// Update profile fields
	employeeprofile.DateOfBirth = updateemployee.DateOfBirth

	employeeprofile.VillageIDOfBirht = updateemployee.VillageIDOfBirht

	employeeprofile.MaterialStatus = updateemployee.MaterialStatus

	employeeprofile.VillageIDCurrentAddress = updateemployee.VillageIDCurrentAddress

	employeeprofile.FamilyPhone = updateemployee.FamilyPhone

	employeeprofile.EducationLevel = updateemployee.EducationLevel

	employeeprofile.ExperienceYear = updateemployee.ExperienceYear

	employeeprofile.PreviousComapy = updateemployee.PreviousComapy

	employeeprofile.BankName = updateemployee.BankName

	employeeprofile.Note = updateemployee.Note

	employeeprofile.PositionLevel = updateemployee.PositionLevel

	if err := config.DB.Save(&employeeprofile).Error; err != nil {

		share.RespondError(c, http.StatusInternalServerError, "Failed to update employee profile")

		return
	}

	share.ResponeSuccess(c, 200, "បុគ្គលិកត្រូវបានកែប្រែ")
}

func ChangeStatusEmployee(c *gin.Context) {
	id := c.Param("id")

	result := config.DB.Model(&models.Employee{}).
		Where("id = ?", id).
		Update("is_active", gorm.Expr("1 - is_active"))

	if result.Error != nil {
		share.RespondError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		share.RespondError(c, http.StatusBadRequest, "មិនមានបុគ្គលិកនេះទេ ឬក៏មិនមានការផ្លាស់ប្តូរ")
		return
	}

	share.ResponeSuccess(c, http.StatusOK, "កែប្រែបានជោគជ័យ")
}
func PromoteEmployee(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Model(&models.Employee{}).Where("id =?", id).Update("is_promote", gorm.Expr("!is_promote"))
	if result.Error != nil {
		share.RespondError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "វាយតម្លៃបានជោគជ៍យ")
}
