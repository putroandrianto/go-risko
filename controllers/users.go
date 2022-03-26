package controllers

import (
	"encoding/json"
	"fmt"
	"go-risko/data"
	"go-risko/models"
	"go-risko/utils"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	var user data.User
	err := d.Decode(&user)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error decoding user: %s", err.Error()))
		return
	}
	if user.Name == "" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Name is required")
		return
	}
	if user.Age <= 0 {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Age is required")
		return
	}
	
	var riskProfile data.RiskProfile
	
	riskAge := 55 - user.Age
	if riskAge >= 30 {
		riskProfile.StockPercent = 72.5
		riskProfile.BondPercent = 21.5
		riskProfile.MMPercent = 100 - (riskProfile.StockPercent + riskProfile.BondPercent)
	} else if riskAge >= 20 {
		riskProfile.StockPercent = 54.5
		riskProfile.BondPercent = 25.5
		riskProfile.MMPercent = 100 - (riskProfile.StockPercent + riskProfile.BondPercent)
	} else {
		riskProfile.StockPercent = 34.5
		riskProfile.BondPercent = 45.5
		riskProfile.MMPercent = 100 - (riskProfile.StockPercent + riskProfile.BondPercent)
	}

	user.RiskProfile = &riskProfile
	result := models.CreateUser(user)
	if result.Error != nil {
		utils.JsonErrorResponse(w, http.StatusInternalServerError, "Internal server error: can't create user")
		return
	}

	utils.JsonSuccessResponse(w, user, "Success create user")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	limit := 10
	offset := 0

	limitQuery := r.URL.Query().Get("limit")
	offsetQuery := r.URL.Query().Get("offset")

	var err error
	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error parsing limit: %s", err.Error()))
			return
		}
	}
	if offsetQuery != "" {
		offset, err = strconv.Atoi(offsetQuery)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error parsing offset: %s", err.Error()))
			return
		}
	}
	
	users := models.GetUsers(limit, offset)
	utils.JsonSuccessResponse(w, users ,"Success get users")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	fmt.Println(vars)

	user, result := models.GetUserById(userId)
	if result.RowsAffected == 1 {
		utils.JsonSuccessResponse(w, user, "Success ger user")
	} else {
		utils.JsonErrorResponse(w, http.StatusNotFound, "User not found")
	}
		
	
}