package models

type Village struct {
	ID         int    `json:"id" gorm:"primarykey"`
	Name       string `json:"name" gorm:"column:name"`
	CommunceID int    `json:"communce_id" gorm:"column:communce_id"`
}
