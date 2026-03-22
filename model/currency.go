package models

type Currency struct {
	ID       int    `json:"id" gorm:"primarykey"`
	Code     string `json:"code" gorm:"column:code"`
	Symbol   string `json:"symbol" gorm:"column:symbol"`
	Name     string `json:"name" gorm:"column:name"`
	Isactive bool   `json:"is_active" gorm:"column:is_active"`
}
type CurrencyResponse struct {
	ID       int    `json:"id" gorm:"primarykey"`
	Code     string `json:"code"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Isactive bool   `json:"is_active"`
}
