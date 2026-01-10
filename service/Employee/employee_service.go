package employeeservice

import (
	"HRbackend/config"
	"HRbackend/helper"
	models "HRbackend/model"
	employeerequstupdate "HRbackend/request/Employee"
	employeeres "HRbackend/response/Employee"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeService interface {
	GetEmployee(filters map[string]string) ([]employeeres.EmployeeResponse, error)
	UpdateEmployee(id int, input employeerequstupdate.EmployeeRequestUpdate, profileImage *multipart.FileHeader, qrImage *multipart.FileHeader, c *gin.Context) error
	ChangeStatusEmployee(id int) error
	PromoteEmployee(id int) error
}

type employeeService struct {
	db *gorm.DB
}

func NewEmployeeService() EmployeeService {
	return &employeeService{
		db: config.DB,
	}
}

func (s *employeeService) GetEmployee(filters map[string]string) ([]employeeres.EmployeeResponse, error) {
	var emplyees []employeeres.EmployeeResponse
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
		Joins("INNER JOIN employee_profiles ON employee_profiles.employee_id = employees.id").
		Joins("INNER JOIN villages AS birth_village ON birth_village.id = employee_profiles.village_id_of_birth").
		Joins("INNER JOIN communces AS birth_communce ON birth_communce.id = birth_village.communce_id").
		Joins("INNER JOIN districts AS birth_district ON birth_district.id = birth_communce.district_id").
		Joins("INNER JOIN provinces AS birth_province ON birth_province.id = birth_district.province_id").
		Joins("INNER JOIN villages AS current_village ON current_village.id = employee_profiles.village_id_current_address").
		Joins("INNER JOIN communces AS current_communce ON current_communce.id = current_village.communce_id").
		Joins("INNER JOIN districts AS current_district ON current_district.id = current_communce.district_id").
		Joins("INNER JOIN provinces AS current_province ON current_province.id = current_district.province_id").
		Joins("INNER JOIN branches ON branches.id = employees.branch_id").
		Joins("INNER JOIN roles ON roles.id = employees.role_id").
		Joins("INNER JOIN employee_shifts ON employee_shifts.employee_id = employees.id AND employee_shifts.is_active = 1").
		Joins("INNER JOIN shifts ON shifts.id = employee_shifts.shift_id").
		Joins("INNER JOIN salaries ON salaries.employee_shift_id = employee_shifts.id AND salaries.is_active = 1").
		Joins("INNER JOIN currencies ON currencies.id = salaries.currency_id")

	if value, ok := filters["branch_id"]; ok && value != "" {
		db = db.Where("employees.branch_id =?", value)
	}
	if value, ok := filters["name"]; ok && value != "" {
		db = db.Where("emplyees.name_en LIKE ? OR employees.name_kh LIKE ?", "%"+value+"%", "%"+value+"%")
	}
	if value, ok := filters["role_id"]; ok && value != "" {
		db = db.Where("employees.role_id =?", value)
	}
	if value, ok := filters["is_active"]; ok && value != "" {
		db = db.Where("employees.is_active =?", value)
	}
	if value, ok := filters["shift_id"]; ok && value != "" {
		db = db.Where("employee_shifts.shift_id =?", value)
	}
	if value, ok := filters["is_promote"]; ok && value != "" {
		db = db.Where("employees.is_promote =?", value)
	}
	if err := db.Order("employees.id DESC").Scan(&emplyees).Error; err != nil {
		return nil, err
	}
	for i := range emplyees {
		emplyees[i].DateOfBirth = helper.FormatDate(emplyees[i].DateOfBirth)
		emplyees[i].HireDate = helper.FormatDate(emplyees[i].HireDate)
		emplyees[i].PromoteDate = helper.FormatDate(emplyees[i].PromoteDate)
	}
	return emplyees, nil
}

func (s *employeeService) UpdateEmployee(
	id int,
	input employeerequstupdate.EmployeeRequestUpdate,
	profileImage *multipart.FileHeader,
	qrImage *multipart.FileHeader,
	c *gin.Context,
) error {

	result := s.db.Model(&models.Employee{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"branch_id":          input.BranchID,
			"name_en":            input.NameEn,
			"name_kh":            input.NameKh,
			"gender":             input.Gender,
			"contact":            input.Contact,
			"national_id_number": input.NationalIDNumber,
			"role_id":            input.RoleID,
			"hire_date":          input.HireDate,
			"promote_date":       input.PromoteDate,
			"type":               input.Type,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("employee not found")
	}
	var employeeprofile models.EmployeeProfile
	if err := s.db.Where("employee_profiles.employee_id =?", id).First(&employeeprofile).Error; err != nil {
		return err
	}
	if profileImage != nil {
		if !helper.ProtectImage(profileImage) {
			return errors.New("រូបភាពមិនត្រូវបានអនុញ្ញាតិបញ្ចូលទេ")
		}
		if employeeprofile.ProfileImage != "" {
			oldProfilePath := filepath.Join("public/profileimage", employeeprofile.ProfileImage)
			if _, err := os.Stat(oldProfilePath); err == nil {
				// os.stat is check file exist or not
				os.Remove(oldProfilePath)

			}
			newProfileImageName, err := s.saveUploadedFile(profileImage, "public/profileimage", c)
			if err != nil {
				return errors.New("កែរូបថ្មីមិនបាន")
			}
			employeeprofile.ProfileImage = newProfileImageName
		}
		if qrImage != nil {
			if !helper.ProtectImage(qrImage) {
				return errors.New("រូបភាពមិនត្រូវបានអនុញ្ញាតិបញ្ចូលទេ")
			}
			if employeeprofile.QrCodeBankAccount != "" {
				oldQRPath := filepath.Join("public/qrcodeimage", employeeprofile.QrCodeBankAccount)
				if _, err := os.Stat(oldQRPath); err == nil {
					os.Remove(oldQRPath)
				}
			}
			newQRImageName, err := s.saveUploadedFile(qrImage, "public/qrcodeimage", c)
			if err != nil {
				return errors.New("កែរូបថ្មីមិនបាន")
			}
			employeeprofile.QrCodeBankAccount = newQRImageName
		}
		employeeprofile.DateOfBirth = input.DateOfBirth
		employeeprofile.VillageIDOfBirht = input.VillageIDOfBirht
		employeeprofile.MaterialStatus = input.MaterialStatus
		employeeprofile.VillageIDCurrentAddress = input.VillageIDCurrentAddress
		employeeprofile.FamilyPhone = input.FamilyPhone
		employeeprofile.EducationLevel = input.EducationLevel
		employeeprofile.ExperienceYear = input.ExperienceYear
		employeeprofile.PreviousComapy = input.PreviousComapy
		employeeprofile.BankName = input.BankName
		employeeprofile.Note = input.Note
		employeeprofile.PositionLevel = input.PositionLevel
		if err := s.db.Save(&employeeprofile).Error; err != nil {
			return errors.New("កែប្រែមិនជោគជ័យ")
		}
	}

	return nil
}
func (s *employeeService) saveUploadedFile(file *multipart.FileHeader, dir string, c *gin.Context) (string, error) {
	// Create directory if not exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	// Generate unique filename
	extension := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
	filePath := filepath.Join(dir, newFileName)

	// Save file
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", err
	}

	return newFileName, nil
}

func (s *employeeService) ChangeStatusEmployee(id int) error {
	result := s.db.Model(&models.Employee{}).Where("id =?", id).Update("is_active", gorm.Expr("1 - is_active"))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("employee not found")
	}
	return nil
}

func (s *employeeService) PromoteEmployee(id int) error {

	result := s.db.Model(&models.Employee{}).Where("id =?", id).Update("is_promote", gorm.Expr("!is_promote"))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("employee not found")
	}
	return nil
}
