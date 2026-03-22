package models

import "time"

type Recieve struct {
	ID           int       `json:"id" gorm:"primarykey"`
	LoanID       int       `json:"loan_id" gorm:"column:loan_id"`
	BranchID     int       `json:"branch_id" gorm:"column:branch_id"`
	RecieveDate  time.Time `json:"receive_date" gorm:"column:receive_date"`
	TotalRecieve float64   `json:"total_receive" gorm:"column:total_receive"`
	RecieveByID  int       `json:"receive_by_id" gorm:"column:receive_by_id"`
	PayrollID    int       `json:"payroll_id" gorm:"column:payroll_id"`
}
