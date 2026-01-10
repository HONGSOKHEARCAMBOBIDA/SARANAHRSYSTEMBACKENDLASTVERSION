package userpart

type UserPartResponse struct {
	ID       int    `json:"id" gorm:"column:id"`
	PartID   int    `json:"part_id"`
	PartName string `json:"part_name"`
}
