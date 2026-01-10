package models

type User struct {
	ID uint `json:"id"`

	BranchID int `json:"branch_id"`

	UserName string `json:"username" gorm:"column:username"`

	Email string `json:"email"`

	Password string `json:"password"`

	Contact string `json:"contact"`

	RoleID int `json:"role_id"`

	EmployeeID int `json:"employee_id"`
}

type UserResponseV2 struct {
	ID      int     `json:"id"`
	NameKh  string  `json:"name_kh"`
	Branch  Branch  `json:"branch"`
	Role    Role    `json:"role"`
	Village Village `json:"village"`
}
