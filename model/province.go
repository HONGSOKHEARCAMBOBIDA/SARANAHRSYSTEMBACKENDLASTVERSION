package models

type Province struct {
	ID   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"column:name"`
}
