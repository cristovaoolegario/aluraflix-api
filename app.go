package main

import (
	"fmt"
	"github.com/cristovaoolegario/aluraflix-api/db"
)

type App struct {
	database db.DatabaseService
}

func (a *App) Initialize(user, password, hostname, dbname string){
	a.database.Server = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", user, password, hostname, dbname)
	a.database.Database = dbname
	a.database.Connect()
}

func (a *App) Run(port string) {
	
}