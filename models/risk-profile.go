package models

type RiskProfile struct {
	Id	uint `json:"id"`
	UserId uint `json:"user_id"`
	MMPercent	float32 `json:"mm_percent"`
	BondPercent	float32 `json:"bond_percent"`
	StockPercent	float32 `json:"stock_percent"`
	User	User `json:"user" gorm:"foreignkey:UserId"`
}