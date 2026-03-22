package models

type District struct {
	ID         int    `json:"id" gorm:"primarykey"`
	Name       string `json:"name" gorm:"column:name"`
	ProvinceID int    `json:"province_id" gorm:"column:province_id"`
}
