package main

import (
	"3-struct/api"
	"3-struct/config"
	"3-struct/services"
	"log"
)

func main() {
	cfg := config.NewConfig()
	storageService := services.NewStorageService("bins.json")
	binService := services.NewBinService()
	jsonBinAPI := api.NewJSONBinAPI(cfg)

	app := services.NewApplicationService(storageService, binService, jsonBinAPI, cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
