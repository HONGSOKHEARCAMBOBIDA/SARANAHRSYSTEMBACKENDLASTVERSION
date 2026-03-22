package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	"HRbackend/helper"

	models "HRbackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateLeave(c *gin.Context) {
	userID, ok := helper.GetUserID(c)
	if !ok {
		share.RespondError(c, http.StatusUnauthorized, "Please Login")
		return
	}
	var input models.LeaveRequestcreate
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	var employeeshift models.EmployeeShift
	if err := config.DB.First(&employeeshift, input.EmployeeShiftID).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, "Employee NOT FOund")
		return
	}
	var employee models.Employee
	if err := config.DB.First(&employee, employeeshift.EmployeeID).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, "Data not found")
		return
	}

	newleave := models.Leave{
		EmployeeShiftID:     input.EmployeeShiftID,
		IsPermission:        input.IsPermission,
		IswithoutPermission: input.IswithoutPermission,
		IsWeeken:            input.IsWeeken,
		StartDate:           input.StartDate,
		EndDate:             input.EndDate,
		LeaveDay:            input.LeaveDay,
		Description:         input.Description,
		Status:              1,
		ApproveById:         input.ApproveById,
		BranchID:            employee.BranchID,
		CreateBy:            userID,
	}
	if err := config.DB.Create(&newleave).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.ResponeSuccess(c, 200, "បង្កេីតបានជោគជ័យ")
}

func GetLeave(c *gin.Context) {
	var leave []models.LeaveResponse

	// --- Query parameters ---
	branchID := c.Query("branch_id")
	employeeID := c.Query("employee_id")
	employeeName := c.Query("employee_name")
	permission := c.Query("permission")
	withoutPermission := c.Query("withoutpermission")
	weekend := c.Query("weekend")
	status := c.Query("status")
	userID, ok := helper.GetUserID(c)
	if !ok {
		share.RespondError(c, http.StatusUnauthorized, "Please Login")
		return
	}
	// --- Base query ---
	db := config.DB.Table("leaves").Select(`
		leaves.id AS id,
		employee_shifts.id AS employee_shift_id,
		employees.id AS employee_id,
		employees.name_en AS employee_name_english,
		employees.name_kh AS employee_name_khmer,
		employees.gender AS gender,
		employees.contact AS contact,
		employees.national_id_number AS national_id_number,
		roles.id AS role_id,
		roles.display_name AS role_name,
		employees.type AS type,
		shifts.id AS shift_id,
		shifts.name AS shift_name,
		shifts.start_time AS start_time,
		shifts.end_time AS end_time,
		branches.id AS branch_id,
		branches.name AS branch_name,
		leaves.is_permission AS is_permission,
		leaves.is_without_permission AS is_without_permission,
		leaves.is_weekend AS is_weekend,
		leaves.start_date AS start_date,
		leaves.end_date AS end_date,
		leaves.leave_days AS leave_days,
		leaves.description AS description,
		leaves.status AS status,
		users.id AS approve_by_id,
		e.name_kh AS approve_by_name
	`).
		Joins("INNER JOIN employee_shifts ON employee_shifts.id = leaves.employee_shift_id").
		Joins("INNER JOIN employees ON employees.id = employee_shifts.employee_id").
		Joins("INNER JOIN shifts ON shifts.id = employee_shifts.shift_id").
		Joins("INNER JOIN roles ON roles.id = employees.role_id").
		Joins("INNER JOIN branches ON branches.id = leaves.branch_id").
		Joins("INNER JOIN users ON users.id = leaves.approve_by_id").
		Joins("INNER JOIN employees e ON e.id = users.employee_id").
		Where("leaves.create_by = ? OR leaves.approve_by_id = ?", userID, userID)

	// --- Filtering ---
	if branchID != "" {
		db = db.Where("leaves.branch_id = ?", branchID)
	}
	if employeeID != "" {
		db = db.Where("employees.id = ?", employeeID)
	}
	if employeeName != "" {
		db = db.Where("(employees.name_en LIKE ? OR employees.name_kh LIKE ?)",
			"%"+employeeName+"%", "%"+employeeName+"%")
	}
	if permission != "" {
		db = db.Where("leaves.is_permission = ?", permission)
	}
	if withoutPermission != "" {
		db = db.Where("leaves.is_without_permission = ?", withoutPermission)
	}
	if weekend != "" {
		db = db.Where("leaves.is_weekend = ?", weekend)
	}
	if status != "" {
		db = db.Where("leaves.status = ?", status)
	}

	// --- Sorting ---
	db = db.Order("leaves.id DESC")

	// --- Execute query ---
	if err := db.Scan(&leave).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	for i := range leave {
		leave[i].Dob = helper.FormatDate(leave[i].Dob)
		leave[i].StartDate = helper.FormatDate(leave[i].StartDate)
		leave[i].EndDate = helper.FormatDate(leave[i].EndDate)
	}

	// --- Success response ---
	share.RespondDate(c, http.StatusOK, leave)
}
func ChangeStatusLeave(c *gin.Context) {
	id := c.Param("id")
	userID, ok := helper.GetUserID(c)
	if !ok {
		share.RespondError(c, http.StatusUnauthorized, "Please Login")
		return
	}
	result := config.DB.Model(&models.Leave{}).Where("id =? AND approve_by_id =? ", id, userID).Update("status", gorm.Expr("1 - status"))
	if result.Error != nil {
		share.RespondError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}
	share.ResponeSuccess(c, 200, "កែប្រែបានជោគជ័យ")
}
func UpdateLeave(c *gin.Context) {
	// --- Get leave ID from URL ---
	id := c.Param("id")

	// --- Parse JSON input ---
	var input models.LeaveRequestcreate
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// --- Find existing leave record ---
	var leave models.Leave
	if err := config.DB.Where("status =?", 1).First(&leave, id).Error; err != nil {
		share.RespondError(c, http.StatusNotFound, "Leave record not found")
		return
	}

	// --- Get related employee data (for branch_id) ---
	var employeeshift models.EmployeeShift
	if err := config.DB.First(&employeeshift, input.EmployeeShiftID).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, "EmployeeShift not found")
		return
	}

	var employee models.Employee
	if err := config.DB.First(&employee, employeeshift.EmployeeID).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, "Employee not found")
		return
	}

	// --- Update fields ---
	leave.EmployeeShiftID = input.EmployeeShiftID
	leave.IsPermission = input.IsPermission
	leave.IswithoutPermission = input.IswithoutPermission
	leave.IsWeeken = input.IsWeeken
	leave.StartDate = input.StartDate
	leave.EndDate = input.EndDate
	leave.LeaveDay = input.LeaveDay
	leave.Description = input.Description
	leave.Status = 1
	leave.ApproveById = input.ApproveById
	leave.BranchID = employee.BranchID

	// --- Save changes ---
	if err := config.DB.Save(&leave).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.ResponeSuccess(c, http.StatusOK, "កែប្រែព័ត៌មានបានជោគជ័យ")
}
