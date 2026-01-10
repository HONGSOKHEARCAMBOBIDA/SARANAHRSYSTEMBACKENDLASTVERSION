package employeeres

type EmployeeResponse struct {
	Id                         int    `json:"id" gorm:"primarykey"`
	BranchShiftID              int    `json:"branch_shift_id"`
	BranchID                   int    `json:"branch_id"`
	BranchName                 string `json:"branch_name"`
	NameEn                     string `json:"name_en"`
	NameKh                     string `json:"name_kh"`
	Gender                     int    `json:"gender"`
	DateOfBirth                string `json:"date_of_birth" gorm:"column:date_of_birth"`
	VillageIDOfBirht           int    `json:"village_id_of_birth" gorm:"column:village_id_of_birth"`
	VillageNameofBirth         string `json:"village_name_of_birth" gorm:"column:village_name_of_birth"`
	CommunceIDofBirth          int    `json:"communce_id_of_birth" gorm:"column:communce_id_of_birth"`
	CommunceNameofBirth        string `json:"communce_name_of_birth" gorm:"column:communce_name_of_birth"`
	DistrictIDofBirth          int    `json:"district_id_of_birth" gorm:"column:district_id_of_birth"`
	DistrictNameofBirth        string `json:"district_name_of_birth" gorm:"column:district_name_of_birth"`
	ProvinceIDofBirth          int    `json:"province_id_of_birth" gorm:"column:province_id_of_birth"`
	ProvinceNameofBirth        string `json:"province_name_of_birth" gorm:"column:province_name_of_birth"`
	MaterialStatus             int    `json:"marital_status" gorm:"column:marital_status"`
	ProfileImage               string `json:"profile_image" gorm:"column:profile_image"`
	VillageIDCurrentAddress    int    `json:"village_id_current_address" gorm:"column:village_id_current_address"`
	VillageNameCurrentAddress  string `json:"village_name_current_address"`
	CommunceIDCurrentAddress   int    `json:"communce_id_current_address"`
	CommunceNameCurrentAdreee  string `json:"communce_name_current_address" gorm:"column:communce_name_current_address"`
	DistrictIDCurrentAddress   int    `json:"district_id_current_address"`
	DistrictNameCurrentAddress string `json:"district_name_current_address"`
	ProvinceIDCurrentAddress   int    `json:"province_id_current_address"`
	ProvinceNameCurrentAddress string `json:"province_name_current_address"`
	FamilyPhone                string `json:"family_phone" gorm:"column:family_phone"`
	EducationLevel             string `json:"education_level" gorm:"column:education_level"`
	ExperienceYear             int    `json:"experience_years" gorm:"column:experience_years"`
	PreviousComapy             string `json:"previous_company" gorm:"column:previous_company"`
	BankName                   string `json:"bank_name" gorm:"column:bank_name"`
	BankAccountNumber          string `json:"bank_account_number" gorm:"column:bank_account_number"`
	QrCodeBankAccount          string `json:"qr_code_bank_account" gorm:"column:qr_code_bank_account"`
	Note                       string `json:"notes" gorm:"column:notes"`
	Contact                    string `json:"contact"`
	NationalIDNumber           string `json:"national_id_number"`
	RoleID                     int    `json:"role_id"`
	RoleName                   string `json:"role_name"`
	PositionLevel              int    `json:"position_level"`
	HireDate                   string `json:"hire_date"`
	PromoteDate                string `json:"promote_date"`
	IsPromote                  bool   `json:"is_promote"`
	Type                       int    `json:"type"`
	ShiftId                    int    `json:"shift_id"`
	ShiftName                  string `json:"shift_name"`
	StartTime                  string `json:"start_time"`
	EndTime                    string `json:"end_time"`
	EmployeeShiftID            int    `json:"employee_shitf_id"`
	SalaryID                   int    `json:"salary_id"`
	BaseSalary                 string `json:"base_salary" gorm:"column:base_salary"`
	WordDay                    int    `json:"worked_day" gorm:"column:worked_day"`
	DailyRate                  string `json:"daily_rate" gorm:"column:daily_rate"`
	IsActive                   bool   `json:"is_active"`
	AssignBranch               int    `json:"assign_branch_id" gorm:"column:assign_branch_id"`
	CurrencyID                 int    `json:"currency_id"`
	CurrencyCode               string `json:"currency_code"`
	CurrencySymbol             string `json:"currency_symbol"`
	CurrencyName               string `json:"currency_name"`
}
