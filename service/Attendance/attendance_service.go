package attendance

import (
	"HRbackend/config"
	"HRbackend/helper"
	models "HRbackend/model"
	attendance "HRbackend/request/Attendance"
	attendanceres "HRbackend/response/Attendance"
	"HRbackend/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AttendanceService interface {
	CheckIn(input attendance.AttendanceLogRequestCreate, userID uint) error
	CheckOut(input attendance.AttendanceLogRequestCreate, userID uint) error
	GetAttendanceLog(filters map[string]string, userID uint) ([]attendanceres.AttendanceResponse, error)
}
type attendanceService struct {
	db *gorm.DB
}

func NewAttendanceService() AttendanceService {
	return &attendanceService{
		db: config.DB,
	}
}
func (as *attendanceService) CheckIn(input attendance.AttendanceLogRequestCreate, userID uint) error {
	var empShift models.EmployeeShift
	if err := as.db.First(&empShift, input.EmployeeShiftID).Error; err != nil {
		return fmt.Errorf("Employee shift NOt found %v", err.Error())
	}
	var shift models.Shift
	if err := as.db.First(&shift, empShift.ShiftID).Error; err != nil {
		return fmt.Errorf("Shift Not Found %v", err.Error())
	}
	var branch models.Branch
	if err := as.db.First(&branch, shift.BranchID).Error; err != nil {
		return fmt.Errorf("Branch Not Found %v", err.Error())
	}
	distance := utils.CalculateDistance(branch.Latitude, branch.Longitude, input.Latitude, input.Longitude)
	isInZone := distance <= branch.Radius
	currentDate := time.Now().Format("2006-01-02")

	var existingLog models.AttendanceLog
	if err := as.db.Where("employee_shift_id = ? AND check_date = ?", input.EmployeeShiftID, currentDate).First(&existingLog).Error; err == nil {
		return fmt.Errorf("you have already checked in today")
	}
	now := time.Now()
	starttime, _ := time.Parse("15:04:05", shift.StartTime)
	isLate := 0
	if now.Hour() > starttime.Hour() || (now.Hour() == starttime.Hour() && now.Minute() > starttime.Minute()) {
		isLate = 1
	}

	newlog := models.AttendanceLog{
		EmployeeShiftID:  input.EmployeeShiftID,
		CheckDate:        currentDate,
		CheckIn:          now.Format("15:04:05"),
		Islate:           isLate,
		BranchID:         branch.ID,
		Status:           1,
		ISZoonCheckIn:    isInZone,
		ISZoonCheckOut:   true,
		LatitudeCheckIn:  input.Latitude,
		LongitudeCheckIn: input.Longitude,
		Notes:            input.Notes,
		CreateBy:         int(userID),
	}
	if err := as.db.Create(&newlog).Error; err != nil {
		return err
	}
	var employee models.Employee
	if err := as.db.Where("id =?", empShift.EmployeeID).First(&empShift).Error; err != nil {
		return fmt.Errorf("employee not found")
	}
	worktime := fmt.Sprintf("%s - %s", shift.StartTime, shift.EndTime)
	mapURL := fmt.Sprintf("https://www.google.com/maps?q=%f,%f",

		input.Latitude,

		input.Latitude,
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
		employee.NameKh,
		employee.Contact,
		branch.Name,
		worktime,
		now.Format("15:04:05"),
		lateText,
		zoneText,
		distance,
		mapURL,
	)
	go helper.SendTelegramMessage(message)
	return nil
	// go as.sendNotification("CHECK IN", uint(empShift.EmployeeID), branch, shift, now, isLate, isInZone, distance, input.Latitude, input.Longitude)
	// return nil

}

func (as *attendanceService) CheckOut(input attendance.AttendanceLogRequestCreate, userID uint) error {
	var empShift models.EmployeeShift
	if err := as.db.First(&empShift, input.EmployeeShiftID).Error; err != nil {
		return fmt.Errorf("Employee Shift Not Found")
	}

	var shift models.Shift
	if err := as.db.First(&shift, empShift.ShiftID).Error; err != nil {
		return fmt.Errorf("Shift Not Found")
	}

	var branch models.Branch
	if err := as.db.First(&branch, shift.BranchID).Error; err != nil {
		return fmt.Errorf("Branch Not Found")
	}

	distance := utils.CalculateDistance(branch.Latitude, branch.Longitude, input.Latitude, input.Longitude)
	isInZone := distance <= branch.Radius

	now := time.Now()

	// Find attendance log by date only

	var log models.AttendanceLog
	if err := as.db.Where("employee_shift_id = ? AND DATE(check_date) = ? AND status = ?", input.EmployeeShiftID, now.Format("2006-01-02"), 1).First(&log).Error; err != nil {
		return fmt.Errorf("Attendance record not found")
	}

	// Early leave calculation
	endTime, _ := time.Parse("15:04:05", shift.EndTime)
	shiftEnd := time.Date(now.Year(), now.Month(), now.Day(), endTime.Hour(), endTime.Minute(), endTime.Second(), 0, now.Location())
	isLeftEarly := 0
	if now.Before(shiftEnd) {
		isLeftEarly = 1
	}

	checkoutTime := now.Format("15:04:05")
	log.CheckOut = &checkoutTime
	log.IsLeftEarly = isLeftEarly
	log.Status = 0
	log.ISZoonCheckOut = isInZone
	log.LatitudeCheckOut = input.Latitude
	log.LongitudeCheckOut = input.Longitude
	log.Notes = input.Notes
	log.CheckDate = now.Format("2006-01-02")

	if err := as.db.Save(&log).Error; err != nil {
		return err
	}

	var employee models.Employee
	if err := as.db.Where("id =?", empShift.EmployeeID).First(&employee).Error; err != nil {
		return fmt.Errorf("employee not found")
	}
	worktime := fmt.Sprintf("%s - %s", shift.StartTime, shift.EndTime)
	mapURL := fmt.Sprintf(
		"https://www.google.com/maps?q=%f,%f",
		input.Latitude,
		input.Longitude,
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
		employee.NameKh,
		employee.Contact,
		branch.Name,
		worktime,
		now.Format("15:04:05"),
		earlyText,
		zoneText,
		distance,
		mapURL,
	)
	go helper.SendTelegramMessage(message)

	// go as.sendNotification("CHECK OUT", uint(empShift.EmployeeID), branch, shift, now, isLeftEarly, isInZone, distance, input.Latitude, input.Longitude)
	return nil
}

func (as *attendanceService) sendNotification(mode string, empID uint, branch models.Branch, shift models.Shift, now time.Time, status int, isInZone bool, distance float64, lat, long float64) {
	var emp models.Employee
	as.db.First(&emp, empID)

	statusIcon := "🟢"
	statusMsg := "ស្កែនទាន់ម៉ោង"
	if status == 1 {
		statusIcon = "🔴"
		statusMsg = "យឺត/ចេញមុន"
	}

	zoneMsg := "📍 ក្នុងតំបន់"
	if !isInZone {
		zoneMsg = "⚠️ ក្រៅតំបន់"
	}

	message := fmt.Sprintf(
		"<b>%s</b>\n👤: %s\n🏢: %s\n🕒: %s\n%s %s\n%s\n📏: %.2fm\n<a href=\"http://maps.google.com/maps?q=%f,%f\">📍 Location</a>",
		mode, emp.NameKh, branch.Name, now.Format("15:04:05"), statusIcon, statusMsg, zoneMsg, distance, lat, long,
	)
	helper.SendTelegramMessage(message)
}

func (as *attendanceService) GetAttendanceLog(filters map[string]string, userID uint) ([]attendanceres.AttendanceResponse, error) {
	var attendance []attendanceres.AttendanceResponse
	db := as.db.Table("attendance_logs").
		Select(`attendance_logs.*, branches.name AS branch_name, employees.name_en, employees.name_kh, roles.display_name AS role_name, shifts.name AS shift_name, shifts.start_time, shifts.end_time`).
		Joins("LEFT JOIN employee_shifts ON employee_shifts.id = attendance_logs.employee_shift_id").
		Joins("LEFT JOIN shifts ON shifts.id = employee_shifts.shift_id").
		Joins("LEFT JOIN employees ON employees.id = employee_shifts.employee_id").
		Joins("LEFT JOIN branches ON branches.id = attendance_logs.branch_id").
		Joins("LEFT JOIN users u ON u.id = ?", userID).
		Joins("LEFT JOIN roles ON roles.id = u.role_id").
		Where("(attendance_logs.create_by = ? OR roles.id IN (1,4,7))", userID)
	if v, ok := filters["branch_id"]; ok && v != "" {
		db = db.Where("attendance_logs.branch_id = ?", v)
	}
	if v, ok := filters["islate"]; ok && v != "" {
		db = db.Where("attendance_logs.is_late = ?", v)
	}
	if v, ok := filters["name"]; ok && v != "" {
		db = db.Where("employees.name_en LIKE ? OR employees.name_kh LIKE ?", "%"+v+"%", "%"+v+"%")
	}

	startDate, sOk := filters["start_date"]
	endDate, eOk := filters["end_date"]
	if sOk && eOk && startDate != "" && endDate != "" {
		db = db.Where("attendance_logs.check_date BETWEEN ? AND ?", startDate, endDate)
	}

	if err := db.Order("attendance_logs.id desc").Scan(&attendance).Error; err != nil {
		return nil, err
	}

	for i := range attendance {
		attendance[i].CheckDate = helper.FormatDate(attendance[i].CheckDate)
	}

	return attendance, nil
}
