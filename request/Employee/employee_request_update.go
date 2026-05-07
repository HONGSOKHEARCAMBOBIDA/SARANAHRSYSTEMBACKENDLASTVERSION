package employeerequstupdate

type EmployeeRequestUpdate struct {
	BranchID                int    `form:"branch_id"`
	NameEn                  string `form:"name_en"`
	NameKh                  string `form:"name_kh"`
	Gender                  int    `form:"gender"`
	Contact                 string `form:"contact"`
	NationalIDNumber        string `form:"national_id_number"`
	RoleID                  int    `form:"role_id"`
	HireDate                string `form:"hire_date"`
	PromoteDate             string `form:"promote_date"`
	Type                    int    `form:"type"`
	DateOfBirth             string `form:"date_of_birth"`
	VillageIDOfBirht        int    `form:"village_id_of_birth"`
	MaterialStatus          int    `form:"marital_status"`
	VillageIDCurrentAddress int    `form:"village_id_current_address"`
	FamilyPhone             string `form:"family_phone"`
	EducationLevel          string `form:"education_level"`
	ExperienceYear          string `form:"experience_years"`
	PreviousComapy          string `form:"previous_company"`
	BankName                string `form:"bank_name" gorm:"column:bank_name"`
	BankAccountNumber       string `form:"bank_account_number"`
	Note                    string `form:"notes"`
	PositionLevel           int    `form:"position_level"`
}
