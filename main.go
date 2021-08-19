package main

import "os"

func main() {

	os.Setenv("AUD", "https://alura-flix-api/")
	os.Setenv("ISS", "https://alura-flix-api.us.auth0.com/")
	a := App{}
	a.Initialize("dev",
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":3000",
		"dev")
}
