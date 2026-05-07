package user

type UserReqInsert struct {
	BranchID int `form:"branch_id" binding:"required"`

	NameEn string `form:"name_en" binding:"required"`

	NameKh string `form:"name_kh" binding:"required"`

	Gender int `form:"gender" binding:"required"`

	Contact string `form:"contact" binding:"required"`

	NationalIDNumber string `form:"national_id_number" binding:"required"`

	RoleID int `form:"role_id" binding:"required"`

	EmployeeID int `form:"employee_id"`

	HireDate string `form:"hire_date" binding:"required"`

	PromoteDate string `form:"promote_date" binding:"required"`

	Type int `form:"type" binding:"required"`

	ShiftID int `form:"shift_id" binding:"required"`

	BaseSalary float64 `form:"base_salary" binding:"required"`

	WorkedDay int `form:"worked_day" binding:"required"`

	EffectTiveDate string `form:"effective_date"`

	DateOfBirth string `form:"date_of_birth" binding:"required"`

	VillageIDOfBirht int `form:"village_id_of_birth" binding:"required"`

	MaterialStatus int `form:"marital_status" binding:"required"`

	VillageIDCurrentAddress int `form:"village_id_current_address" binding:"required"`

	FamilyPhone string `form:"family_phone" binding:"required"`

	EducationLevel string `form:"education_level" binding:"required"`

	ExperienceYear string `form:"experience_years" binding:"required"`

	PreviousComapy string `form:"previous_company" binding:"required"`

	BankName string `form:"bank_name" binding:"required"`

	BankAccountNumber string `form:"bank_account_number" binding:"required"`

	PositionLevel int `form:"position_level" binding:"required"`

	Note string `form:"note"`

	CurrencyID int `form:"currency_id" binding:"required"`

	PartIDs []int `form:"part_ids" binding:"required"`
}
