package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	"HRbackend/helper"
	models "HRbackend/model"
	attendance "HRbackend/request/Attendance"
	attendanceres "HRbackend/response/Attendance"
	"HRbackend/utils"
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 📍 Check In Function
func CheckIn(c *gin.Context) {

	var req attendance.AttendanceLogRequestCreate

	userID, ok := helper.GetUserID(c)

	if !ok {

		share.RespondError(c, http.StatusUnauthorized, "Please Login")

		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	// Step 1: Find EmployeeShift
	var empShift models.EmployeeShift

	if err := config.DB.First(&empShift, req.EmployeeShiftID).Error; err != nil {

		share.RespondError(c, http.StatusNotFound, "Employee Shift Not Found")

		return
	}

	// Step 2: Find shift
	var shift models.Shift

	if err := config.DB.First(&shift, empShift.ShiftID).Error; err != nil {

		share.RespondError(c, http.StatusNotFound, "Shift Not Found")

		return
	}

	// Step 3: Find branch
	var branch models.Branch

	if err := config.DB.First(&branch, shift.BranchID).Error; err != nil {

		share.RespondError(c, http.StatusNotFound, "Branch Not Found")

		return
	}

	// Step 4: Calculate distance
	distance := utils.CalculateDistance(branch.Latitude, branch.Longitude, req.Latitude, req.Longitude)

	currentDate := time.Now().Format("2006-01-02")

	// Step 5: Prevent duplicate check-in
	var existingLog models.AttendanceLog

	if err := config.DB.Where("employee_shift_id = ? AND check_date = ?", req.EmployeeShiftID, currentDate).
		First(&existingLog).Error; err == nil {

		share.RespondError(c, http.StatusNotFound, "You have already checked in today")

		return
	}

	// Step 6: Check lateness
	layout := "15:04:05"

	now := time.Now()

	startTime, _ := time.Parse(layout, shift.StartTime)

	isLate := 0

	if now.After(startTime) {

		isLate = 1

	}

	// Step 7: Create log (inside or outside zone)
	isInZone := distance <= branch.Radius

	log := models.AttendanceLog{

		EmployeeShiftID: req.EmployeeShiftID,

		CheckDate: currentDate,

		CheckIn: now.Format("15:04:05"),

		Islate: isLate,

		BranchID: branch.ID,

		Status: 1,

		ISZoonCheckIn: isInZone, // ✅ true if inside zone, false if outside

		ISZoonCheckOut: true,

		LatitudeCheckIn: req.Latitude,

		LongitudeCheckIn: req.Longitude,

		Notes: req.Notes,

		CreateBy: userID,
	}

	if err := config.DB.Create(&log).Error; err != nil {

		share.RespondError(c, http.StatusInternalServerError, err.Error())

		return
	}

	var singleemployee models.Employee

	if err := config.DB.Where("id = ?", empShift.EmployeeID).First(&singleemployee).Error; err != nil {

		share.RespondError(c, http.StatusNotFound, err.Error())

		return
	}

	workTime := fmt.Sprintf("%s - %s", shift.StartTime, shift.EndTime)

	mapURL := fmt.Sprintf("https://www.google.com/maps?q=%f,%f",

		req.Latitude,

		req.Longitude,
	)

	lateText := "⏰ ស្កែនទាន់ម៉ោង"

	if isLate == 1 {

		lateText = "🔴 ចូលធ្វេីការយឺត"

	}

	zoneText := "📍 ស្កែនក្នុងតំបន់ក្រុមហ៊ុន"

	if !isInZone {

		zoneText = "⚠️ ស្កែនក្រៅតំបន់ក្រុមហ៊ុន"

	}

	message := fmt.Sprintf(

		"🟢 <b>CHECK IN</b>\n\n"+
			"👤 ឈ្មោះ: %s\n"+
			"📲 លេខទូរសព្ទ: %s\n"+
			"🏢 សាខា: %s\n"+
			"🕒 ម៉ោងធ្វើការ: %s\n"+
			"🕒 Check-in: %s\n"+
			"%s\n"+
			"%s\n"+
			"📏 Distance: %.2f m\n"+
			"🗺 <a href=\"%s\">មេីលទីតាំងស្កែន</a>",
		singleemployee.NameKh,
		singleemployee.Contact,
		branch.Name,
		workTime,
		now.Format("15:04:05"),
		lateText,
		zoneText,
		distance,
		mapURL,
	)

	go helper.SendTelegramMessage(message)

	share.ResponeSuccess(c, 200, "Check In Success")

}

func CheckOut(c *gin.Context) {
	var req attendance.AttendanceLogRequestCreate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Step 1: Get employee shift
	var empShift models.EmployeeShift
	if err := config.DB.First(&empShift, req.EmployeeShiftID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee shift not found"})
		return
	}

	// Step 2: Get shift
	var shift models.Shift
	if err := config.DB.First(&shift, empShift.ShiftID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shift not found"})
		return
	}

	// Step 3: Get branch
	var branch models.Branch
	if err := config.DB.First(&branch, shift.BranchID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
		return
	}

	// Step 4: Calculate distance
	distance := utils.CalculateDistance(branch.Latitude, branch.Longitude, req.Latitude, req.Longitude)
	isInZone := distance <= branch.Radius // ✅ true if inside, false if outside

	// Step 5: Check early leave
	layout := "15:04:05"
	now := time.Now()
	endTimeParsed, _ := time.Parse(layout, shift.EndTime)

	endTime := time.Date(
		now.Year(), now.Month(), now.Day(),
		endTimeParsed.Hour(), endTimeParsed.Minute(), endTimeParsed.Second(),
		0, now.Location(),
	)

	isLeftEarly := 0
	if now.Before(endTime) {
		isLeftEarly = 1
	}

	// Step 6: Find today's attendance record (status = 1 = checked in)
	var log models.AttendanceLog
	if err := config.DB.Where("employee_shift_id = ? AND check_date = ? AND status = ?", req.EmployeeShiftID, now.Format("2006-01-02"), 1).
		First(&log).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attendance record not found"})
		return
	}

	// Step 7: Update check-out info
	checkout := now.Format("15:04:05")
	log.CheckOut = &checkout
	log.CheckDate = now.Format("2006-01-02")

	log.IsLeftEarly = isLeftEarly
	log.Status = 0                // 0 = checked out
	log.ISZoonCheckOut = isInZone // ✅ mark whether user was in zone
	log.LatitudeCheckOut = req.Latitude
	log.LongitudeCheckOut = req.Longitude
	log.Notes = req.Notes

	if err := config.DB.Save(&log).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var singleemployee models.Employee
	if err := config.DB.
		Where("id = ?", empShift.EmployeeID).
		First(&singleemployee).Error; err != nil {
		share.RespondError(c, http.StatusNotFound, err.Error())
		return
	}
	workTime := fmt.Sprintf("%s - %s", shift.StartTime, shift.EndTime)
	mapURL := fmt.Sprintf(
		"https://www.google.com/maps?q=%f,%f",
		req.Latitude,
		req.Longitude,
	)
	earlyText := "⏰ ស្កែនត្រូវម៉ោង"

	if isLeftEarly == 1 {
		earlyText = "🔴 ចេញមុនម៉ោងកំណត់"
	}
	zoneText := "📍 ស្កែនក្នុងតំបន់ក្រុមហ៊ុន"
	if !isInZone {
		zoneText = "⚠️ ស្កែនក្រៅតំបន់ក្រុមហ៊ុន"
	}
	message := fmt.Sprintf(
		"🟢 <b>CHECK OUT</b>\n\n"+
			"👤 ឈ្មោះ: %s\n"+

			"📲 លេខទូរសព្ទ: %s\n"+
			"🏢 សាខា: %s\n"+
			"🕒 ម៉ោងធ្វើការ: %s\n"+
			"🕒 Check-out: %s\n"+
			"%s\n"+
			"%s\n"+
			"📏 Distance: %.2f m\n"+
			"🗺 <a href=\"%s\">មេីលទីតាំងស្កែន</a>",
		singleemployee.NameKh,
		singleemployee.Contact,
		branch.Name,
		workTime,
		now.Format("15:04:05"),
		earlyText,
		zoneText,
		distance,
		mapURL,
	)
	go helper.SendTelegramMessage(message)
	share.ResponeSuccess(c, http.StatusOK, "Check-out successful")
}

func GetAttendanceLog(c *gin.Context) {
	var attendance []attendanceres.AttendanceResponse

	branchID := c.Query("branch_id")
	isLate := c.Query("islate")
	isLeftEarly := c.Query("isleftearly")
	name := c.Query("name")
	startDate := c.Query("start_date") // e.g., 2025-10-01
	endDate := c.Query("end_date")     // e.g., 2025-10-31
	userID, ok := helper.GetUserID(c)
	if !ok {
		share.RespondError(c, http.StatusUnauthorized, "Please Login")
		return
	}

	db := config.DB.Table("attendance_logs").
		Select(`
			attendance_logs.id AS id,
			attendance_logs.employee_shift_id AS employee_shift_id,
			attendance_logs.check_date AS check_date,
			attendance_logs.check_in AS check_in,
			attendance_logs.check_out AS check_out,
			attendance_logs.is_late AS is_late,
			attendance_logs.is_left_early AS is_left_early,
			attendance_logs.is_zoon_check_in AS is_zoon_check_in,
			attendance_logs.is_zoon_check_out AS is_zoon_check_out,
			attendance_logs.latitude_check_in AS latitude_check_in,
			attendance_logs.longitude_check_in AS longitude_check_in,
			attendance_logs.latitude_check_out AS latitude_check_out,
			attendance_logs.longitude_check_out AS longitude_check_out,
			attendance_logs.notes AS notes,
			branches.id AS branch_id,
			branches.name AS branch_name,
			attendance_logs.status AS status,
			employees.name_en AS name_en,
			employees.name_kh AS name_kh,
			employee_profiles.profile_image AS employee_profile,
			roles.id AS role_id,
			roles.display_name AS role_name,
			employees.type AS type,
			shifts.id AS shift_id,
			shifts.name AS shift_name,
			shifts.start_time AS start_time,
			shifts.end_time AS end_time
		`).Where("(attendance_logs.create_by = ? OR roles.id IN (1,4,7))", userID).
		Joins("LEFT JOIN employee_shifts ON employee_shifts.id = attendance_logs.employee_shift_id").
		Joins("LEFT JOIN shifts ON shifts.id = employee_shifts.shift_id").
		Joins("LEFT JOIN employees ON employees.id = employee_shifts.employee_id").
		Joins("LEFT JOIN employee_profiles ON employee_profiles.employee_id = employees.id").
		Joins("LEFT JOIN branches ON branches.id = attendance_logs.branch_id").
		Joins("LEFT JOIN users u ON u.id =?", userID).
		Joins("LEFT JOIN roles ON roles.id = u.role_id")

	// Optional filters
	if branchID != "" {
		db = db.Where("attendance_logs.branch_id = ?", branchID)
	}
	if isLate != "" {
		db = db.Where("attendance_logs.is_late = ?", isLate)
	}
	if isLeftEarly != "" {
		db = db.Where("attendance_logs.is_left_early = ?", isLeftEarly)
	}
	if name != "" {
		db = db.Where("employees.name_en LIKE ? OR employees.name_kh LIKE ?", "%"+name+"%", "%"+name+"%")
	}

	// ✅ Date range filter (new)
	if startDate != "" && endDate != "" {
		db = db.Where("attendance_logs.check_date BETWEEN ? AND ?", startDate, endDate)
	} else if startDate != "" {
		db = db.Where("attendance_logs.check_date >= ?", startDate)
	} else if endDate != "" {
		db = db.Where("attendance_logs.check_date <= ?", endDate)
	}
	db = db.Order("attendance_logs.id desc")

	if err := db.Scan(&attendance).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	for i := range attendance {
		attendance[i].CheckDate = helper.FormatDate(attendance[i].CheckDate)
	}

	share.RespondDate(c, http.StatusOK, attendance)
}
