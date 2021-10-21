package main

import "os"

func main() {

	a := initApp()

	a.Run(os.Getenv("PORT"),
		  os.Getenv("ENV"))
}
