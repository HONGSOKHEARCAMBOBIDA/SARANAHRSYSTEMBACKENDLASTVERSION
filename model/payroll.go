package models

type Payroll struct {
	ID                            int     `json:"id"`
	SalaryID                      int     `json:"salary_id" gorm:"column:salary_id"`
	PayrollYear                   int     `json:"payroll_year" gorm:"column:payroll_year"`
	PayrollMonth                  int     `json:"payrollmonth" gorm:"column:payrollmonth"`
	GrossSalary                   float64 `json:"grosssalary" gorm:"column:grosssalary"`
	TotalLate                     int     `json:"total_late" gorm:"column:total_late"`
	LatePenalty                   int     `json:"latepenalty" gorm:"column:latepenalty"`
	ToalEarlyexit                 int     `json:"total_earlyexit" gorm:"column:total_earlyexit"`
	TotalExitpenalty              int     `json:"totalexitpenalty" gorm:"column:totalexitpenalty"`
	LeaveWithPermission           int     `json:"leave_with_permission" gorm:"column:leave_with_permission"`
	PeanaltyLeaveWithPermission   float64 `json:"penalty_leave_with_permission" gorm:"column:penalty_leave_with_permission"`
	LeaveWithoutPermission        int     `json:"leave_without_permission" gorm:"column:leave_without_permission"`
	PenaltyLeaveWithoutPermission float64 `json:"penalty_leave_without_permission" gorm:"column:penalty_leave_without_permission"`
	LeaveWeekend                  int     `json:"leave_weekend" gorm:"column:leave_weekend"`
	PenaltyLeaveWeekend           float64 `json:"penalty_leave_weekend" gorm:"column:penalty_leave_weekend"`
	LoanDeduction                 float64 `json:"loanDeduction" gorm:"column:loanDeduction"`
	IsAttendanceBonus             int     `json:"is_attendance_bonus" gorm:"column:is_attendance_bonus"`
	BonusType                     string  `json:"bonus_type" gorm:"column:bonus_type"`
	BonusAmount                   float64 `json:"bonus_amount" gorm:"column:bonus_amount"`
	TotalDeduction                float64 `json:"totalDeductions" gorm:"column:totalDeductions"`
	NetSalary                     float64 `json:"netsalary" gorm:"column:netsalary"`
	Status                        int     `json:"status" gorm:"column:status"`
	BranchId                      int     `json:"branch_id" gorm:"column:branch_id"`
	CurrencyID                    int     `json:"currency_id" gorm:"column:currency_id"`
	ExchangeRate                  float64 `json:"exchange_rates" gorm:"column:exchange_rates"`
}

type PayrollRequestCreate struct {
	SalaryID                      int     `json:"salary_id" gorm:"column:salary_id"`
	PayrollMonth                  int     `json:"payrollmonth" gorm:"column:payrollmonth"`
	GrossSalary                   float64 `json:"notdeduction" gorm:"column:notdeduction"`
	TotalLate                     int     `json:"total_late" gorm:"column:total_late"`
	LatePenalty                   int     `json:"penaltylate" gorm:"column:penaltylate"`
	ToalEarlyexit                 int     `json:"total_earlyexit" gorm:"column:total_earlyexit"`
	TotalExitpenalty              int     `json:"totalexitpenalty" gorm:"column:totalexitpenalty"`
	LeaveWithPermission           int     `json:"leave_with_permission" gorm:"column:leave_with_permission"`
	PeanaltyLeaveWithPermission   float64 `json:"penalty_leave_with_permission" gorm:"column:penalty_leave_with_permission"`
	LeaveWithoutPermission        int     `json:"leave_without_permission" gorm:"column:leave_without_permission"`
	PenaltyLeaveWithoutPermission float64 `json:"penalty_leave_without_permission" gorm:"column:penalty_leave_without_permission"`
	LeaveWeekend                  int     `json:"leave_weekend" gorm:"column:leave_weekend"`
	PenaltyLeaveWeekend           float64 `json:"penalty_leave_weekend" gorm:"column:penalty_leave_weekend"`
	LoanID                        int     `json:"loan_id"`
	LoanDeduction                 float64 `json:"loanDeduction" gorm:"column:loanDeduction"`
	IsAttendanceBonus             int     `json:"is_attendance_bonus" gorm:"column:is_attendance_bonus"`
	BonusType                     string  `json:"bonus_type" gorm:"column:bonus_type"`
	BonusAmount                   float64 `json:"bonus_amount" gorm:"column:bonus_amount"`
	TotalDeduction                float64 `json:"totalDeductions" gorm:"column:totalDeductions"`
	NetSalary                     float64 `json:"netsalary" gorm:"column:netsalary"`
	Status                        int     `json:"status" gorm:"column:status"`
	BranchId                      int     `json:"branch_id" gorm:"column:branch_id"`
	CurrencyID                    int     `json:"currency_id" gorm:"column:currency_id"`
	ExchangeRate                  float64 `json:"exchange_rates" gorm:"column:exchange_rates"`
}
type PayResponse struct {
	ID                            int    `json:"id" gorm:"column:id"`
	SalaryID                      int    `json:"salary_id" gorm:"column:salary_id"`
	EmployeeShiftID               int    `json:"employee_shift_id" gorm:"column:employee_shift_id"`
	BaseSalary                    string `json:"base_salary" gorm:"column:base_salary"`
	WorkDay                       int    `json:"worked_day" gorm:"column:worked_day"`
	DailyRate                     string `json:"daily_rate" gorm:"column:daily_rate"`
	EmployeeID                    int    `json:"employee_id" gorm:"column:employee_id"`
	NameEn                        string `json:"name_en" gorm:"column:name_en"`
	NameKh                        string `json:"name_kh" gorm:"column:name_kh"`
	Gender                        int    `json:"gender" gorm:"column:gender"`
	RoleID                        int    `json:"role_id" gorm:"column:role_id"`
	RoleName                      string `json:"role_name" gorm:"column:role_name"`
	ShiftID                       int    `json:"shift_id" gorm:"column:shift_id"`
	ShiftName                     string `json:"shift_name" gorm:"column:shift_name"`
	StartTime                     string `json:"start_time" gorm:"column:start_time"`
	EndTime                       string `json:"end_time" gorm:"column:end_time"`
	Payrollyear                   int    `json:"payroll_year" gorm:"column:payroll_year"`
	PayrollMonth                  int    `json:"payrollmonth" gorm:"column:payrollmonth"`
	GrossSalary                   string `json:"grosssalary" gorm:"column:grosssalary"`
	TotalLate                     int    `json:"total_late" gorm:"column:total_late"`
	LatePenalty                   int    `json:"latepenalty" gorm:"column:latepenalty"`
	TotalEarlyExit                int    `json:"total_earlyexit" gorm:"column:total_earlyexit"`
	TotalEarlyPenaltyExit         int    `json:"totalexitpenalty" gorm:"column:totalexitpenalty"`
	LeaveWithPermission           int    `json:"leave_with_permission" gorm:"column:leave_with_permission"`
	PeanaltyLeaveWithPermission   int    `json:"penalty_leave_with_permission" gorm:"column:penalty_leave_with_permission"`
	LeaveWithoutPermission        int    `json:"leave_without_permission" gorm:"column:leave_without_permission"`
	PenaltyLeaveWithoutPermission string `json:"penalty_leave_without_permission" gorm:"column:penalty_leave_without_permission"`
	LeaveWeekend                  int    `json:"leave_weekend" gorm:"column:leave_weekend"`
	PenaltyLeaveWeekend           int    `json:"penalty_leave_weekend" gorm:"column:penalty_leave_weekend"`
	LoanDeduction                 string `json:"loanDeduction" gorm:"column:loanDeduction"`
	IsAttendanceBonus             int    `json:"is_attendance_bonus" gorm:"column:is_attendance_bonus"`
	BonusType                     string `json:"bonus_type" gorm:"column:bonus_type"`
	BonusAmount                   int    `json:"bonus_amount" gorm:"column:bonus_amount"`
	TotalDeduction                string `json:"totalDeductions" gorm:"column:totalDeductions"`
	NetSalary                     string `json:"netsalary" gorm:"column:netsalary"`
	Status                        int    `json:"status" gorm:"column:status"`
	BranchID                      int    `json:"branch_id" gorm:"column:branch_id"`
	BranchName                    string `json:"branch_name" gorm:"column:branch_name"`
	CurrencyID                    int    `json:"currency_id" gorm:"column:currency_id"`
	CurrencyCode                  string `json:"currency_code" gorm:"column:currency_code"`
	CurrencySymbol                string `json:"currency_symbol" gorm:"column:currency_symbol"`
	CurrencyName                  string `json:"currency_name" gorm:"column:currency_name"`
	BaseCurrencyID                int    `json:"base_currency_id"`
	BaseCurrencyCode              string `json:"base_currency_code"`
	BaseCurrencySymbol            string `json:"base_currency_symbol"`
	ExchangeRate                  string `json:"exchange_rate" gorm:"column:exchange_rate"`
}
