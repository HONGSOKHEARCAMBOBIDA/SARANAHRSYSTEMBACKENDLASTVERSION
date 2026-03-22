package models

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey" json:"role_id"`
	PermissionID uint `gorm:"primaryKey" json:"permission_id"`
}

func (RolePermission) TableName() string {
	return "role_has_permissions"
}

type CreateRolePermissionInput struct {
	RoleID        int   `json:"role_id" binding:"required"`
	PermissionIDs []int `json:"permission_ids" binding:"required"`
}

type DeleteRolePermissionsInput struct {
	RoleID        int   `json:"role_id" binding:"required"`
	PermissionIDs []int `json:"permission_ids" binding:"required"`
}

type PermissionWithAssignedRole struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Assigned    bool   `json:"assigned"`
}
