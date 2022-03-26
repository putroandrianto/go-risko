package app

import (
	"go-risko/controllers"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.Index)
	// router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	// router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")

	return router
}