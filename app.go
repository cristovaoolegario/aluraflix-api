package main

import (
	"fmt"
	"github.com/cristovaoolegario/aluraflix-api/db"
	"github.com/cristovaoolegario/aluraflix-api/router"
	"github.com/gorilla/mux"
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

func (a *App) Run(port string) {
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, a.router))
}
