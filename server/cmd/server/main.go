package main

import (
	"GoGameV3/internal/handler"
	"GoGameV3/internal/service"
	"GoGameV3/pkg/libreoffice"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	{
		v1 := router.Group("/v1")

		v1.POST("/files/upload", handler.FileUpload)
		v1.POST("/files/:fileId/convert", handler.ConvertFileById)
		v1.GET("/files/:convertedFileId", handler.GetInfoConvertFileById)
		v1.GET("/files/:convertedFileId/download", handler.DownloadFileById)
	}

	conventor := libreoffice.LibreOffice{}
	service := service.NewService(conventor)
	service.HealthCheck()
	router.Run(":8080")
}
