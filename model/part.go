package models

type Part struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name"`
}

type PartResquest struct {
	Name string `json:"name"`
}
