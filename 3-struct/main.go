package main

import (
	"3-struct/api"
	"3-struct/config"
	"3-struct/services"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found: " + err.Error())
	}

	cfg := config.NewConfig()
	storageService := services.NewStorageService("bins.json")
	binService := services.NewBinService()
	jsonBinAPI := api.NewJSONBinAPI(cfg)

	app := services.NewApplicationService(storageService, binService, jsonBinAPI, cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
