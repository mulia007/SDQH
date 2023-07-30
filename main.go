package main

import (
	"log"
	"net/http"
	"sdqh/controllers/authcontroller"
	"sdqh/controllers/datacontroller"
	"sdqh/middlewares"
	"sdqh/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/data", datacontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8000", r))
}
