package app

import (
	"fmt"
	"github.com/cristovaoolegario/aluraflix-api/internal/pkg/storage/bson/db/services"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type App struct {
	router   *mux.Router
	database services.DatabaseService
}

func ProvideApp(router mux.Router, db services.DatabaseService) App {
	return App{&router, db}
}

func (a *App) Run(port, env string) {
	fmt.Println("Server running in port:", port)
	stringedPort := fmt.Sprintf(":%s", port)
	if env == "dev" {
		corsWrapper := cors.New(cors.Options{
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
		})
		log.Fatal(http.ListenAndServe(stringedPort, corsWrapper.Handler(a.router)))
	} else {
		log.Fatal(http.ListenAndServe(stringedPort, a.router))
	}
}
