package main

import (
	"GoGameV3/internal/handler"
	"GoGameV3/internal/router"
	"GoGameV3/internal/service"
	"GoGameV3/middleware/cors"
	"GoGameV3/pkg/libreoffice"
)

func main() {
	converter := libreoffice.NewConverter()
	service := service.NewService(converter)
	handler := handler.NewHandler(service)
	router := router.SetupRouter(handler)
	router.Use(cors.CORSMidlleware())
	router.Run(":8080")

}
