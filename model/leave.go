package models

type Leave struct {
	ID                  int    `json:"id" gorm:"primarykey"`
	EmployeeShiftID     int    `json:"employee_shift_id" gorm:"column:employee_shift_id"`
	IsPermission        int    `json:"is_permission" gorm:"column:is_permission"`
	IswithoutPermission int    `json:"is_without_permission" gorm:"column:is_without_permission"`
	IsWeeken            int    `json:"is_weekend" gorm:"column:is_weekend"`
	StartDate           string `json:"start_date" gorm:"column:start_date"`
	EndDate             string `json:"end_date" gorm:"column:end_date"`
	LeaveDay            int    `json:"leave_days" gorm:"column:leave_days"`
	Description         string `json:"description" gorm:"column:description"`
	Status              int    `json:"status" gorm:"column:status"`
	ApproveById         int    `json:"approve_by_id" gorm:"column:approve_by_id"`
	BranchID            int    `json:"branch_id" gorm:"column:branch_id"`
	CreateBy            int    `json:"create_by"`
}

type LeaveRequestcreate struct {
	EmployeeShiftID     int    `json:"employee_shift_id"`
	IsPermission        int    `json:"is_permission"`
	IswithoutPermission int    `json:"is_without_permission"`
	IsWeeken            int    `json:"is_weekend"`
	StartDate           string `json:"start_date"`
	EndDate             string `json:"end_date"`
	LeaveDay            int    `json:"leave_days"`
	Description         string `json:"description"`
	Status              int    `json:"status"`
	ApproveById         int    `json:"approve_by_id"`
}
type LeaveResponse struct {
	ID                  int    `json:"id"`
	EmployeeShiftID     int    `json:"employee_shift_id"`
	EmployeeID          int    `json:"employee_id"`
	EmployeeNameEnglish string `json:"employee_name_english"`
	EmployeeNameKhmer   string `json:"employee_name_khmer"`

	Gender              int    `json:"gender"`
	Dob                 string `json:"dob"`
	Contact             string `json:"contact"`
	NationalIDNumber    string `json:"national_id_number"`
	RoleID              int    `json:"role_id"`
	RoleName            string `json:"role_name"`
	EmployeeType        int    `json:"type" gorm:"column:type"`
	ShiftID             int    `json:"shift_id"`
	ShiftName           string `json:"shift_name"`
	StartTime           string `json:"start_time"`
	EndTime             string `json:"end_time"`
	BranchID            int    `json:"branch_id"`
	BranchName          string `json:"branch_name"`
	IsPermission        int    `json:"is_permission"`
	IswithoutPermission int    `json:"is_without_permission" gorm:"column:is_without_permission"`
	IsWeeken            int    `json:"is_weekend" gorm:"column:is_weekend"`
	StartDate           string `json:"start_date"`
	EndDate             string `json:"end_date"`
	LeaveDay            int    `json:"leave_days" gorm:"column:leave_days"`
	Description         string `json:"description"`
	Status              int    `json:"status"`
	ApproveById         int    `json:"approve_by_id"`
	ApproveByName       string `json:"approve_by_name"`
}
