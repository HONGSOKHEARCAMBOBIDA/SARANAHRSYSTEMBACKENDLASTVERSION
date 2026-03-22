package models

type EmployeeProfile struct {
	ID int `json:"id" gorm:"primarykey"`

	EmployeeID int `json:"employee_id" gorm:"column:employee_id"`

	DateOfBirth string `json:"date_of_birth" gorm:"column:date_of_birth"`

	VillageIDOfBirht int `json:"village_id_of_birth" gorm:"column:village_id_of_birth"`

	MaterialStatus int `json:"marital_status" gorm:"column:marital_status"`

	ProfileImage string `json:"profile_image" gorm:"column:profile_image"`

	VillageIDCurrentAddress int `json:"village_id_current_address" gorm:"column:village_id_current_address"`

	FamilyPhone string `json:"family_phone" gorm:"column:family_phone"`

	EducationLevel string `json:"education_level" gorm:"column:education_level"`

	ExperienceYear int `json:"experience_years" gorm:"column:experience_years"`

	PreviousComapy string `json:"previous_company" gorm:"column:previous_company"`

	BankName string `json:"bank_name" gorm:"column:bank_name"`

	BankAccountNumber string `json:"bank_account_number" gorm:"column:bank_account_number"`

	QrCodeBankAccount string `json:"qr_code_bank_account" gorm:"column:qr_code_bank_account"`

	Note string `json:"notes" gorm:"column:notes"`

	PositionLevel int `form:"position_level"`
}
