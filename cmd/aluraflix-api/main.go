// @title Aluraflix API
// @version 1.0
// @description This is a sample service for managing videos and categories
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email cristovaoolegario@gmail.com
// @license.name MIT
// @license.url https://spdx.org/licenses/MIT.html
// @host https://cristovao-aluraflix-api.herokuapp.com
// @BasePath /api/v1
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
