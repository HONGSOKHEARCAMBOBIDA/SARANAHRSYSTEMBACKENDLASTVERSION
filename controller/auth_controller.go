package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	"HRbackend/helper"
	models "HRbackend/model"
	user "HRbackend/request/User"
	userresponse "HRbackend/response/User"
	userpart "HRbackend/response/UserPart"
	"HRbackend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {

	var input user.UserReqInsert

	if err := c.ShouldBind(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	hashedPassword := utils.HashPassword(input.Password)

	file, err := c.FormFile("profileimage")

	if err != nil {

		share.RespondError(c, http.StatusBadRequest, "profileimage is required")

		return
	}

	if !helper.ProtectImage(file) {

		share.RespondError(c, http.StatusBadRequest, "Only PNG and JPG images are allowed for Profile Image")

	}

	profileimageDir := "public/profileimage"

	if _, err := os.Stat(profileimageDir); os.IsNotExist(err) {

		os.MkdirAll(profileimageDir, os.ModePerm)

	}

	extension := filepath.Ext(file.Filename)

	profileimageName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)

	profileimagePath := filepath.Join(profileimageDir, profileimageName)

	if err := c.SaveUploadedFile(file, profileimagePath); err != nil {

		share.RespondError(c, http.StatusInternalServerError, "Faild to insert image")

		return
	}

	fileqrcode, err := c.FormFile("qrcodeimage")

	if err != nil {

		share.RespondError(c, http.StatusBadRequest, "image qr is required")

		return
	}

	if !helper.ProtectImage(fileqrcode) {

		share.RespondError(c, http.StatusBadRequest, "Only PNG and JPG images are allowed for QR Code")

		return

	}

	qrcodeimageDir := "public/qrcodeimage"

	if _, err := os.Stat(qrcodeimageDir); os.IsNotExist(err) {

		os.MkdirAll(qrcodeimageDir, os.ModePerm)

	}

	extensionqr := filepath.Ext(fileqrcode.Filename)

	qrimageName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extensionqr)

	qrimagepath := filepath.Join(qrcodeimageDir, qrimageName)

	if err := c.SaveUploadedFile(fileqrcode, qrimagepath); err != nil {

		share.RespondError(c, http.StatusInternalServerError, "Fail to insert image qr")
		return
	}

	tx := config.DB.Begin()

	if tx.Error != nil {

		share.RespondError(c, 500, tx.Error.Error())

		return
	}

	emp := models.Employee{

		BranchID: input.BranchID,

		NameEn: input.NameEn,

		NameKh: input.NameKh,

		Gender: input.Gender,

		Contact: input.Contact,

		NationalIDNumber: input.NationalIDNumber,

		RoleID: input.RoleID,

		HireDate: input.HireDate,

		PromoteDate: input.PromoteDate,

		Type: input.Type,
	}

	if err := tx.Create(&emp).Error; err != nil {

		tx.Rollback()

		share.RespondError(c, 500, err.Error())

		return
	}

	emprofile := models.EmployeeProfile{

		EmployeeID: int(emp.ID),

		DateOfBirth: input.DateOfBirth,

		VillageIDOfBirht: input.VillageIDOfBirht,

		MaterialStatus: input.MaterialStatus,

		ProfileImage: profileimageName,

		VillageIDCurrentAddress: input.VillageIDCurrentAddress,

		FamilyPhone: input.FamilyPhone,

		EducationLevel: input.EducationLevel,

		ExperienceYear: input.ExperienceYear,

		PreviousComapy: input.PreviousComapy,

		BankName: input.BankName,

		BankAccountNumber: input.BankAccountNumber,

		QrCodeBankAccount: qrimageName,

		Note: input.Note,

		PositionLevel: input.PositionLevel,
	}

	if err := tx.Create(&emprofile).Error; err != nil {

		tx.Rollback()

		share.RespondError(c, 500, err.Error())

		return
	}

	user := models.User{

		BranchID: input.BranchID,

		UserName: input.UserName,

		Email: input.Email,

		Password: hashedPassword, // ✅ Use hashed

		Contact: input.Contact,

		RoleID: input.RoleID,

		EmployeeID: int(emp.ID),
	}

	if err := tx.Create(&user).Error; err != nil {

		tx.Rollback()

		share.RespondError(c, 500, tx.Error.Error())

		return
	}

	if len(input.PartIDs) > 0 {

		for _, part := range input.PartIDs {

			if err := tx.Create(&models.UserPart{

				UserID: int(user.ID),

				PartID: part,
			}).Error; err != nil {

				tx.Rollback()

				share.RespondError(c, http.StatusInternalServerError, err.Error())

				return
			}
		}
	}

	employeeshift := models.EmployeeShift{

		EmployeeID: int(emp.ID),

		ShiftID: input.ShiftID,
	}

	if err := tx.Create(&employeeshift).Error; err != nil {

		tx.Rollback()

		share.RespondError(c, 500, err.Error())

		return
	}

	salary := models.Salary{

		EmployeeShiftID: employeeshift.ID,

		BaseSalary: input.BaseSalary,

		WorkedDay: input.WorkedDay,

		DailyRate: float64(input.BaseSalary) / float64(input.WorkedDay),

		EffectTiveDate: time.Now(), // field type must be time.Time

		CurrencyID: input.CurrencyID,
	}

	if err := tx.Create(&salary).Error; err != nil {

		tx.Rollback()

		share.RespondError(c, 500, err.Error())

		return
	}

	tx.Commit()

	share.ResponeSuccess(c, 201, "ចុះឈ្មោះបានជោគជ័យ")

}

func ChangeStatusUser(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {

		share.RespondError(c, http.StatusNotFound, err.Error())

		return
	}

	result := config.DB.Model(&models.User{}).Where("id =?", id).Update("is_active", gorm.Expr("1 - is_active"))

	if result.Error != nil {

		share.RespondError(c, 500, result.Error.Error())

		return

	}
	if result.RowsAffected == 0 {

		share.RespondError(c, http.StatusNotFound, result.Error.Error())

		return
	}

	share.RespondDate(c, 200, result)
}

func UpdateUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		share.RespondError(c, http.StatusBadRequest, "Invalid ID")

		return
	}

	var updateuser user.UserReqUpdate

	if err := c.ShouldBindJSON(&updateuser); err != nil {

		share.RespondError(c, http.StatusBadRequest, err.Error())

		return
	}

	tx := config.DB.Begin()

	if tx.Error != nil {

		share.RespondError(c, http.StatusInternalServerError, tx.Error.Error())

		return
	}

	// 1️⃣ Update user
	result := tx.Model(&models.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{

			"branch_id": updateuser.BranchID,
			"username":  updateuser.UserName,
			"email":     updateuser.Email,
			"contact":   updateuser.Contact,
			"role_id":   updateuser.RoleID,
		})

	if result.Error != nil {

		tx.Rollback()

		share.RespondError(c, http.StatusInternalServerError, result.Error.Error())

		return
	}

	// 2️⃣ Delete old parts
	if err := tx.Where("user_id = ?", id).
		Delete(&models.UserPart{}).Error; err != nil {

		tx.Rollback()

		share.RespondError(c, http.StatusInternalServerError, err.Error())

		return
	}

	// 3️⃣ Insert new parts
	for _, partID := range updateuser.PartIDs {

		if err := tx.Create(&models.UserPart{
			UserID: id,
			PartID: partID,
		}).Error; err != nil {

			tx.Rollback()

			share.RespondError(c, http.StatusInternalServerError, err.Error())

			return
		}
	}

	tx.Commit()

	share.ResponeSuccess(c, http.StatusOK, "អ្នកប្រើប្រាស់ត្រូវបានកែប្រែ")
}

func GetUser(c *gin.Context) {

	var users []userresponse.UserResponse

	branchID := c.Query("branch_id")

	name := c.Query("name")

	roleID := c.Query("role_id")

	isActive := c.Query("is_active")

	cacheKey := "users:all" +
		":branch:" + branchID +
		":role:" + roleID +
		":name:" + name +
		":active:" + isActive

	cached, err := utils.Redis.Get(utils.Ctx, cacheKey).Result()
	if err == nil {
		// Cache hit
		var users []userresponse.UserResponse
		if err := json.Unmarshal([]byte(cached), &users); err == nil {
			share.RespondDate(c, http.StatusOK, users)
			return
		}
		// If unmarshal fails → ignore cache and query DB
	}

	db := config.DB.Table("users").Select(`

        users.id AS id,

        branches.id AS branch_id,

        branches.name AS branch_name,

        employees.name_en AS name_en,

        employees.name_kh AS name_kh,

        users.username AS username,

        users.email AS email,

        employees.gender AS gender,

        users.contact AS contact,

        employees.national_id_number AS national_id_number,
		
        users.is_active AS is_active,

        roles.id AS role_id,

        roles.display_name AS role_name
    `).
		Joins("INNER JOIN branches ON branches.id = users.branch_id").
		Joins("INNER JOIN roles ON roles.id = users.role_id").
		Joins("INNER JOIN employees ON employees.id = users.employee_id")

	if branchID != "" {

		db = db.Where("users.branch_id = ?", branchID)
	}

	if name != "" {

		db = db.Where("employees.name_en LIKE ? OR employees.name_kh LIKE ?", "%"+name+"%", "%"+name+"%")

	}

	if roleID != "" {

		db = db.Where("users.role_id = ?", roleID)

	}

	if isActive != "" {

		db = db.Where("users.is_active = ?", isActive)

	}

	if err := db.Order("users.id DESC").Scan(&users).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(users) > 0 {
		userIDs := make([]int, 0, len(users))
		for _, u := range users {
			userIDs = append(userIDs, u.Id)
		}
		// Struct សម្រាប់ទទួលទិន្នន័យពី SQL
		type partRow struct {
			UserID   int
			ID       int
			PartID   int
			PartName string
		}

		var rows []partRow
		if err := config.DB.Raw(`
		
		SELECT up.user_id,up.id,p.id AS part_id,p.name AS part_name FROM user_parts up

		JOIN parts p ON p.id = up.part_id WHERE up.user_id IN (?)

		`, userIDs).Scan(&rows).Error; err != nil {
			share.RespondError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//បង្កើត Map ដើម្បីរៀបចំ Part តាម User

		partMap := make(map[int][]userpart.UserPartResponse)

		for _, r := range rows {
			partMap[r.UserID] = append(partMap[r.UserID], userpart.UserPartResponse{
				ID:       r.ID,
				PartID:   r.PartID,
				PartName: r.PartName,
			})
		}

		for i := range users {
			users[i].UserPartResponse = partMap[users[i].Id]
		}

	}
	data, _ := json.Marshal(users)
	utils.Redis.Set(utils.Ctx, cacheKey, data, 30*time.Second)

	share.RespondDate(c, http.StatusOK, users)
}
