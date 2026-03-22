package models

type SummaryforPayroll struct {
	SalaryID                      int     `json:"salary_id"`
	EmployeeShiftID               int     `json:"employee_shift_id"`
	BaseSalary                    float32 `json:"base_salary"`
	Workday                       int     `json:"worked_day" gorm:"column:worked_day"`
	DailyRate                     float32 `json:"daily_rate" gorm:"column:daily_rate"`
	ShiftID                       int     `json:"shift_id"`
	ShiftName                     string  `json:"shift_name"`
	StartTime                     string  `json:"start_time"`
	EndTime                       string  `json:"end_time"`
	BranchID                      int     `json:"branch_id"`
	BranchName                    string  `json:"branch_name"`
	EmployeeID                    int     `json:"employee_id"`
	NameEn                        string  `json:"name_en"`
	NameKh                        string  `json:"name_kh"`
	Gender                        int     `json:"gender"`
	GenderText                    string  `json:"gender_text" gorm:"-"`
	Dob                           string  `json:"dob"`
	Contact                       string  `json:"contact"`
	RoleID                        int     `json:"role_id"`
	RoleName                      string  `json:"role_name"`
	Type                          int     `json:"type"`
	TypeText                      string  `json:"type_text"`
	HireDate                      string  `json:"hire_date"`
	TotalLate                     int     `json:"total_late"`
	PenaltyLate                   float32 `json:"penaltylate" gorm:"column:penaltylate"`
	TotalEarlyExit                int     `json:"total_earlyexit" gorm:"column:total_earlyexit"`
	TotalExitpenalty              float32 `json:"totalexitpenalty" gorm:"column:totalexitpenalty"`
	LeaveWithPermission           int     `json:"leave_with_permission"`
	PeanaltyLeaveWithPermission   float32 `json:"penalty_leave_with_permission" gorm:"column:penalty_leave_with_permission"`
	LeaveWithoutPermission        int     `json:"leave_without_permission"`
	PenaltyLeaveWithoutPermission float32 `json:"penalty_leave_without_permission"`
	LeaveWeekend                  int     `json:"leave_weekend"`
	PenaltyLeaveWeekend           float32 `json:"penalty_leave_weekend"`
	LoanID                        int     `json:"loan_id"`
	LoanAmount                    float32 `json:"loan_amount"`
	RemainingBalance              float32 `json:"remaining_balance" gorm:"column:remaining_balance"`
	AttendanceCount               int     `json:"attendancecount" gorm:"column:attendancecount"`
	NotDeduction                  float32 `json:"notdeduction" gorm:"column:notdeduction"`
	TotalDeduction                float32 `json:"totalDeductions" gorm:"column:totalDeductions"`
	NetSalary                     float32 `json:"netsalary" gorm:"column:netsalary"`
	IsBonusAttendance             int     `json:"is_bonus_attendanace" gorm:"column:is_bonus_attendance"`
	CurrencyID                    int     `json:"currency_id"`
	CurrencyCode                  string  `json:"currency_code"`
	CurrencySymbol                string  `json:"currency_symbol"`
	CurrencyName                  string  `json:"currency_name"`
	TargetCurrencyID              int     `json:"target_currency_id"`
	TargetCurrencyCode            string  `json:"target_currency_code"`
	TargetCurrencySymbol          string  `json:"target_currency_symbol"`
	TargetCurrencyName            string  `json:"target_currency_name"`
	CurrencyPaireID               int     `json:"currency_pair_id" gorm:"column:currency_pair_id"`
	ExchangeRate                  float64 `josn:"exchange_rate"`
}
