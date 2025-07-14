package main

import (
	"3-struct/services"
	"log"
)

func main() {
	storageService := services.NewStorageService("bins.json")
	binService := services.NewBinService()

	app := services.NewApplicationService(storageService, binService)

	if err := app.Run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
