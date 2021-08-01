package router

import "github.com/gorilla/mux"

func Router() *mux.Router{
	r := mux.Router{}
	r.HandleFunc("/api/v1/videos", GetAll).Methods("GET")
	r.HandleFunc("/api/v1/videos/{id}", GetByID).Methods("GET")
	r.HandleFunc("/api/v1/videos", Create).Methods("POST")
	r.HandleFunc("/api/v1/videos/{id}", Update).Methods("PUT")
	return &r
}