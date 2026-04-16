package main

import (
	"GoGameV3/internal/handler"
	"GoGameV3/internal/service"
	"GoGameV3/middleware/cors"
	"GoGameV3/pkg/libreoffice"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.CORSMidlleware())

	var mockFlag = flag.Bool("mock", false, "mock server")
	flag.Parse()

	//NOTE for front-end reason we
	if *mockFlag {
		fmt.Println("Mock server is activating...")
		router.POST("/files/upload", handler.MockFileUpload)
		router.POST("/files/:fileId/convert", handler.MockConvertFileById)
		router.GET("/files/:convertedFileId/download", handler.MockDownloadFileById)
	} else {
		fmt.Println("Dev server is activating...")
		router.POST("/files/upload", handler.FileUpload)
		router.POST("/files/:fileId/convert", handler.ConvertFileById)
		router.GET("/files/:convertedFileId/download", handler.DownloadFileById)
	}

	conventor := libreoffice.LibreOffice{}
	service := service.NewService(conventor)
	service.HealthCheck()
	service.FileCheck("/home/cyber/Documents/GoGameV3/server/files/mock.docx")
	router.Run(":8080")
}
