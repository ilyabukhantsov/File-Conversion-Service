package main

import (
	"log"
	"os"

	"GoGameV3/internal/handler"
	"GoGameV3/internal/router"
	"GoGameV3/internal/service"
	"GoGameV3/middleware/cors"
	"GoGameV3/pkg/libreoffice"
)

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	dataDir := getenv("DATA_DIR", "./data")
	port := getenv("PORT", "8080")

	converter := libreoffice.NewConverter()

	svc, err := service.NewService(converter, dataDir)
	if err != nil {
		log.Fatalf("init service: %v", err)
	}

	h := handler.NewHandler(svc)
	r := router.SetupRouter(h, cors.CORSMidlleware())

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("run server: %v", err)
	}
}
