package app

import (
	"go-risko/controllers"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/users", controllers.GetUsers).Methods("OPTIONS", "GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("OPTIONS", "POST")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("OPTIONS", "GET")

	return router
}