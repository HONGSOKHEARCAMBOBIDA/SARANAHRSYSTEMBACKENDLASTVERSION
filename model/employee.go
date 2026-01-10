package models

type Employee struct {
	ID               uint   `gorm:"primarykey" form:"id"`
	BranchID         int    `json:"branch_id"`
	NameEn           string `json:"name_en"`
	NameKh           string `json:"name_kh"`
	Gender           int    `json:"gender"`
	Contact          string `json:"contact"`
	NationalIDNumber string `json:"national_id_number"`
	RoleID           int    `json:"role_id"`
	HireDate         string `json:"hire_date"`
	PromoteDate      string `json:"promote_date"`
	Type             int    `json:"type"`
}
