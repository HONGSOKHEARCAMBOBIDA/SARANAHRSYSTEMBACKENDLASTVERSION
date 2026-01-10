package currencypairgo

type CurrencyPairRequest struct {
	BaseCurrencyID   int `json:"base_currency_id" gorm:"column:base_currency_id"`
	TargetCurrencyID int `json:"target_currency_id" gorm:"column:target_currency_id"`
}
