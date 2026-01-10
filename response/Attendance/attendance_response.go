package attendanceres

type AttendanceResponse struct {
	ID                int    `json:"id"`
	EmployeeShiftID   int    `json:"employee_shift_id "`
	CheckDate         string `json:"check_date"`
	CheckIn           string `json:"check_in"`
	CheckOut          string `json:"check_out"`
	IsLate            int    `json:"is_late"`
	IsLeftEarly       int    `json:"is_left_early"`
	BranchID          int    `json:"branch_id"`
	BranchName        string `json:"branch_name"`
	Status            int    `json:"status"`
	Nameen            string `json:"name_en" gorm:"column:name_en"`
	NameKh            string `json:"name_kh" gorm:"column:name_kh"`
	RoleID            int    `json:"role_id"`
	RoleName          string `json:"role_name"`
	Type              int    `json:"type" gorm:"column:type"`
	ShiftID           int    `json:"shift_id"`
	ShiftName         string `json:"shift_name"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
	ISZoonCheckIn     bool   `json:"is_zoon_check_in" gorm:"column:is_zoon_check_in"`
	ISZoonCheckOut    bool   `json:"is_zoon_check_out" gorm:"column:is_zoon_check_out"`
	LatitudeCheckIn   string `json:"latitude_check_in" gorm:"column:latitude_check_in"`
	LongitudeCheckIn  string `json:"longitude_check_in" gorm:"column:longitude_check_in"`
	LatitudeCheckOut  string `json:"latitude_check_out" gorm:"column:latitude_check_out"`
	LongitudeCheckOut string `json:"longitude_check_out" gorm:"column:longitude_check_out"`
	Notes             string `json:"notes"`
}
