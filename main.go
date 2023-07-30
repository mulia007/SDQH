package main

import (
	"log"
	"net/http"
	"sdqh/controllers/admincontroller"
	"sdqh/controllers/authcontroller"
	"sdqh/controllers/datacontroller"
	"sdqh/controllers/kegiatancontroller"
	"sdqh/controllers/registerkegiatancontroller"
	"sdqh/controllers/transaksicontroller"
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
	r.HandleFunc("/admin/login", admincontroller.Login).Methods("POST")
	r.HandleFunc("/admin/register", admincontroller.Register).Methods("POST")
	r.HandleFunc("/admin/logout", admincontroller.Logout).Methods("GET")
	r.HandleFunc("/admin/registerkegiatan", registerkegiatancontroller.CreateKegiatan).Methods("POST")
	r.HandleFunc("/getkegiatan", kegiatancontroller.Index).Methods("GET")
	r.HandleFunc("/getkegiatan/detail", kegiatancontroller.Detail).Methods("GET")
	r.HandleFunc("/transaksi/create", transaksicontroller.CreateTransaksi).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/data", datacontroller.Index).Methods("GET")

	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8000", r))
}
