package attendance

type AttendanceLogRequestCreate struct {
	EmployeeShiftID int `json:"employee_shift_id" binding:"required"`

	Latitude float64 `json:"latitude" binding:"required"`

	Longitude float64 `json:"longitude" binding:"required"`

	Notes string `json:"notes"`
}
