package router

import "github.com/gorilla/mux"

func Router() *mux.Router{
	r := mux.Router{}
	r.HandleFunc("/api/v1/videos", GetAll).Methods("GET")
	r.HandleFunc("/api/v1/videos/{id}", GetByID).Methods("GET")
	return &r
}