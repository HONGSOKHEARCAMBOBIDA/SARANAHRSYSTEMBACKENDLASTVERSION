package models

type Shift struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	BranchID  int    `json:"branch_id" gorm:"column:branch_id"`
	IsActive  int    `json:"is_active" gorm:"column:is_active"`
}

type ShiftRequest struct {
	Name      string `json:"name"`
	StartTime string `json:"start_time"` // e.g. "08:00 AM"
	EndTime   string `json:"end_time"`   // e.g. "05:00 PM"
	BranchID  int    `json:"branch_id"`
}
type ShiftResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	BranchID  int    `json:"branch_id"`
	IsActive  int    `json:"is_active"`
}
