package models

type Loan struct {
	ID              int     `json:"id" gorm:"primarykey"`
	EmployeeID      int     `json:"employee_id" gorm:"column:employee_id"`
	BranchID        int     `json:"branch_id" gorm:"column:branch_id"`
	CurrencyID      int     `json:"currency_id"`
	LoanAmount      float64 `json:"loan_amount" gorm:"column:loan_amount"`
	RemainingAmount float64 `json:"remaining_balance" gorm:"column:remaining_balance"`
	Status          int     `json:"status" gorm:"status"`
}

type LoanRequestCreate struct {
	EmployeeID int     `json:"employee_id" gorm:"column:employee_id"`
	CurrencyID int     `json:"currency_id"`
	LoanAmount float64 `json:"loan_amount" gorm:"loan_amount"`
}
type LoanRequestupdate struct {
	EmployeeID      int     `json:"employee_id" gorm:"column:employee_id"`
	LoanAmount      float64 `json:"loan_amount" gorm:"column:loan_amount"`
	CurrencyID      int     `json:"currency_id"`
	RemainingAmount float64 `json:"remaining_balance" gorm:"column:remaining_balance"`
}
type LoanResponse struct {
	ID               int     `json:"id"`
	EmployeeID       int     `json:"employee_id"`
	CurrencyID       int     `json:"currency_id"`
	CurrencyName     string  `json:"currency_name"`
	CurrencyCode     string  `json:"currency_code"`
	CurrencySymbol   string  `json:"currency_symbol"`
	EmployeeName     string  `json:"employee_name"`
	BranchID         int     `json:"branch_id"`
	BranchName       string  `json:"branch_name"`
	LoanAmount       float64 `json:"loan_amount"`
	RemainingBalance float64 `json:"remaining_balance"`
	Status           int     `json:"status"`
}
