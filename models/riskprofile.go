package models

type RiskProfile struct {
	UserId 			uint 	`json:"-"`
	MMPercent		float32 `json:"mm_percent"`
	BondPercent		float32 `json:"bond_percent"`
	StockPercent	float32 `json:"stock_percent"`
}