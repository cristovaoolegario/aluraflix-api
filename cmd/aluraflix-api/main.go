// @title Aluraflix API
// @version 1.0
// @description This is a sample service for managing videos and categories

// @contact.name API Support
// @contact.email cristovaoolegario@gmail.com
// @license.name MIT
// @license.url https://spdx.org/licenses/MIT.html

// @host localhost:3000
//cristovao-aluraflix-api.herokuapp.com
// @Schemes https http
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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
