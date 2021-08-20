package main

import "os"

func main() {

	a := App{}
	a.Initialize(os.Getenv("ENV"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_NAME"))

	a.Run(os.Getenv("PORT"),
		  os.Getenv("ENV"))
}
