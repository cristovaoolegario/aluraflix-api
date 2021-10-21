package main

import (
	"fmt"
	"github.com/cristovaoolegario/aluraflix-api/db"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type App struct {
	router   *mux.Router
	database db.DatabaseService
}

func ProvideApp(router mux.Router, db db.DatabaseService) App {
	return App{&router, db}
}

func (a *App) Run(port, env string) {
	fmt.Println("Server running in port:", port)
	if env == "dev" {
		corsWrapper := cors.New(cors.Options{
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
		})
		log.Fatal(http.ListenAndServe(port, corsWrapper.Handler(a.router)))
	} else {
		log.Fatal(http.ListenAndServe(port, a.router))
	}
}
