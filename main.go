package main

import (
	"github.com/joho/godotenv"
	"os"
)

func main() {

	godotenv.Load()
	a := InitApp()

	a.Run(os.Getenv("PORT"),
		  os.Getenv("ENV"))
}
