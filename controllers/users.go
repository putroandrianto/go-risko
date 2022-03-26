package controllers

import (
	"net/http"

	"go-risko/connections"
	"go-risko/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	type NewUser struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}

	db := connections.DB
	user := new(models.User)
	var i = 55
	switch {
		case i-user.Age >= 30 :
			StockPercent := 72.5
			BondPercent := 21.5
			MMPercent := 100 - (StockPercent + BondPercent)

			riskProfile := models.RiskProfile{
				UserId:     user.Id,
				StockPercent: float32(StockPercent),
				BondPercent:  float32(BondPercent),
				MMPercent:    float32(MMPercent),	
			}
			db.Create(&riskProfile)

		case i-user.Age >= 20 :
			StockPercent := 54.5
			BondPercent := 25.5
			MMPercent := 100 - (StockPercent + BondPercent)

			riskProfile := models.RiskProfile{
				UserId:     user.Id,
				StockPercent: float32(StockPercent),
				BondPercent:  float32(BondPercent),
				MMPercent:    float32(MMPercent),
			}
			db.Create(&riskProfile)

		case i-user.Age < 20 :
			StockPercent := 34.5
			BondPercent := 45.5
			MMPercent := 100 - (StockPercent + BondPercent)

			riskProfile := models.RiskProfile{
				UserId:     user.Id,
				StockPercent: float32(StockPercent),
				BondPercent:  float32(BondPercent),
				MMPercent:    float32(MMPercent),
			}
			db.Create(&riskProfile)
	}
}