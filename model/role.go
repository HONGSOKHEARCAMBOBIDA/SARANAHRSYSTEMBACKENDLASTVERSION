package models

type Role struct {
	ID          uint         `gorm:"primarykey" json:"id"`
	Name        string       `json:"name"`
	DisPlayName string       `json:"display_name" gorm:"column:display_name"`
	IsActive    bool         `json:"is_active"`
	Permissions []Permission `gorm:"many2many:role_has_permissions"`
}
type RoleResponse struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Name        string `json:"name"`
	DisPlayName string `json:"display_name" gorm:"column:display_name"`
	IsActive    bool   `json:"is_active"`
}
type RoleRequestUpdate struct {
	Name        string `json:"name"`
	DisPlayName string `json:"display_name" gorm:"column:display_name"`
	IsActive    bool   `json:"is_active"`
}
type RoleRequestCreate struct {
	Name        string `json:"name"`
	DisPlayName string `json:"display_name" gorm:"column:display_name"`
	IsActive    bool   `json:"is_active"`
}
