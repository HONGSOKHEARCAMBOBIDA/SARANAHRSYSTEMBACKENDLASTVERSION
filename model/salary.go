package models

import "time"

type Salary struct {
	ID              int       `json:"id" gorm:"primarykey"`
	EmployeeShiftID int       `json:"employee_shift_id" gorm:"column:employee_shift_id"`
	BaseSalary      float64   `json:"base_salary" gorm:"column:base_salary"`
	WorkedDay       int       `json:"worked_day" gorm:"column:worked_day"`
	DailyRate       float64   `json:"daily_rate" gorm:"column:daily_rate"`
	EffectTiveDate  time.Time `json:"effective_date" gorm:"column:effective_date"`
	ExpireDate      time.Time `json:"expire_date" gorm:"expire_date"`
	CurrencyID      int       `json:"currency_id" gorm:"column:currency_id"`
}

type SalaryRequestCreate struct {
	EmployeeShiftID int       `json:"employee_shift_id" gorm:"column:employee_shift_id"`
	BaseSalary      float64   `json:"base_salary" gorm:"column:base_salary"`
	WorkedDay       int       `json:"worked_day" gorm:"column:worked_day"`
	DailyRate       float64   `json:"daily_rate" gorm:"column:daily_rate"`
	EffectTiveDate  time.Time `json:"effective_date" gorm:"column:effective_date"`
	ExpireDate      time.Time `json:"expire_date" gorm:"expire_date"`
	CurrencyID      int       `json:"currency_id" gorm:"column:currency_id"`
}

type SalaryRequestUpdate struct {
	BaseSalary float64 `json:"base_salary"`
	Workday    int     `json:"worked_day"`
	DailyRate  float64 `json:"daily_rate"`
	CurrencyID int     `json:"currency_id"`
}
