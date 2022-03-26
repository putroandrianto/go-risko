package controllers

import (
	"encoding/json"
	"fmt"
	"go-risko/models"
	"go-risko/utils"

	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	var user models.User
	err := d.Decode(&user)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error decoding user: %s", err.Error()))
		return
	}
	if user.Name =="" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Name is required")
		return
	}
	if user.Age == 0 {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Age is required")
		return
	}
	
	var riskProfile models.RiskProfile
	
	compareAge := 55 - user.Age
	if compareAge >= 30 {
		riskProfile.StockPercent = 72.5
		riskProfile.BondPercent = 21.5
		riskProfile.MMPercent = 100 - (riskProfile.StockPercent + riskProfile.BondPercent)
	} else if compareAge >= 20 {
		riskProfile.StockPercent = 54.5
		riskProfile.BondPercent = 25.5
		riskProfile.MMPercent = 100 - (riskProfile.StockPercent + riskProfile.BondPercent)
	} else {
		riskProfile.StockPercent = 34.5
		riskProfile.BondPercent = 45.5
		riskProfile.MMPercent = 100 - (riskProfile.StockPercent + riskProfile.BondPercent)
	}

	utils.JsonSuccessResponse(w, http.StatusCreated, "User created", user)
}