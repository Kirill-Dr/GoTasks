package main

import (
	"3-struct/config"
	"3-struct/services"
	"log"
)

func main() {
	cfg := config.NewConfig()
	storageService := services.NewStorageService("bins.json")
	binService := services.NewBinService()

	app := services.NewApplicationService(storageService, binService, cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
