package models

type EmployeeShift struct {
	ID             int `json:"id" gorm:"primarykey"`
	EmployeeID     int `json:"employee_id" gorm:"column:employee_id"`
	ShiftID        int `json:"shift_id" gorm:"column:shift_id"`
	AssignBranchID int `json:"assign_branch_id" gorm:"column:assign_branch_id"`
}
type EmployeeShiftRequestUpdate struct {
	EmployeeID     int `json:"employee_id" gorm:"column:employee_id"`
	ShiftID        int `json:"shift_id" gorm:"column:shift_id"`
	AssignBranchID int `json:"assign_branch_id" gorm:"column:assign_branch_id"`
}
type EmployeeShftRequestCreate struct {
	EmployeeID     int `json:"employee_id" gorm:"column:employee_id"`
	ShiftID        int `json:"shift_id" gorm:"column:shift_id"`
	AssignBranchID int `json:"assign_branch_id" gorm:"column:assign_branch_id"`
}
type EmployeeShiftResponse struct {
	ID        int    `json:"id"`
	ShiftID   int    `json:"shift_id"`
	ShiftName string `json:"shift_name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
