package user

type UserReqUpdate struct {
	BranchID int `json:"branch_id"`

	UserName string `json:"username"`

	Email string `json:"email"`

	Contact string `json:"contact"`

	RoleID int `json:"role_id"`

	PartIDs []int `json:"part_ids"`
}
