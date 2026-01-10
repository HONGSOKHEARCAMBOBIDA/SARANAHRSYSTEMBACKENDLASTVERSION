package user

import userpart "HRbackend/response/UserPart"

type UserResponse struct {
	Id int `json:"id" gorm:"primarykey"`

	BranchID int `json:"branch_id"`

	BranchName string `json:"branch_name"`

	NameEn string `json:"name_en"`

	NameKh string `json:"name_kh"`

	UserName string `json:"username" gorm:"column:username"`

	Email string `json:"email"`

	Gender int `json:"gender"`

	Contact string `json:"contact"`

	NationalIDNumber string `json:"national_id_number"`

	RoleID int `json:"role_id"`

	RoleName string `json:"role_name"`

	IsActive bool `json:"is_active"`

	UserPartResponse []userpart.UserPartResponse `json:"parts" gorm:"-"`
}
