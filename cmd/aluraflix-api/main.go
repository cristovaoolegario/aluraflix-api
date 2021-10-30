package main

import (
	"github.com/cristovaoolegario/aluraflix-api/internal/app"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	godotenv.Load()
	a := app.InitApp()

	a.Run(os.Getenv("PORT"),
		  os.Getenv("ENV"))
}
