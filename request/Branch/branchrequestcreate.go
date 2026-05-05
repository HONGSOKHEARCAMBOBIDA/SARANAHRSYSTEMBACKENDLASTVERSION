package branch

type BranchRequestCreate struct {
	Name      string   `json:"name" gorm:"column:name"`
	Latitude  *float64 `json:"latitude" gorm:"column:latitude"`
	Longitude *float64 `json:"longitude" gorm:"column:longitude"`
	Radius    float64  `json:"radius" gorm:"column:radius"`
}
