package main

import "os"

func main() {

	a := InitApp()

	a.Run(os.Getenv("PORT"),
		  os.Getenv("ENV"))
}
