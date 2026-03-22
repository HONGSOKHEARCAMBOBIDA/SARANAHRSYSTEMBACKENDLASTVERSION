package models

type Communce struct {
	ID         int    `json:"id" gorm:"primarykey"`
	Name       string `json:"name" gorm:"column:name"`
	DistrictID int    `json:"district_id" gorm:"column:district_id"`
}
