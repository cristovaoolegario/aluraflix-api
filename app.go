package main

import (
	"fmt"
	"github.com/cristovaoolegario/aluraflix-api/db"
	"github.com/cristovaoolegario/aluraflix-api/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type App struct {
	router   *mux.Router
	database db.DatabaseService
}

func (a *App) Initialize(user, password, hostname, dbname string) {
	a.database.Server = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", user, password, hostname, dbname)
	a.database.Database = dbname
	a.database.Connect()
	a.router = router.Router()
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
