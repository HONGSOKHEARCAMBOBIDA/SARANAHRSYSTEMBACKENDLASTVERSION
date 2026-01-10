package models

type Branch struct {
	ID        int     `json:"id" gorm:"primarykey"`
	Name      string  `json:"name" gorm:"column:name"`
	Latitude  float64 `json:"latitude" gorm:"column:latitude"`
	Longitude float64 `json:"longitude" gorm:"column:longitude"`
	Radius    float64 `json:"radius" gorm:"column:radius"`
	IsActive  int     `json:"is_active"`
}
