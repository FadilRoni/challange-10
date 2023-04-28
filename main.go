package main

import (
	"challange-10/app"
	"challange-10/config"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	err := config.InitPostgres()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApp()
}
