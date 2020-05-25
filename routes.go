package main

import (
	"log"

	"github.com/gorilla/mux"
)

// AddRoutes :
func AddRoutes(router *mux.Router) {

	log.Println("Loading routes ... ")

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/add", addQafiya).Methods("POST")

	log.Println("Routes are loaded :)")
}
