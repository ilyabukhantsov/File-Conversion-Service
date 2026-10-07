package router

import (
	"github.com/gin-gonic/gin"
)

// Интерфейс живет здесь. Он описывает, ЧТО роутеру нужно от хендлера.
type Handler interface {
	UploadFile(c *gin.Context)
	ConvertFile(c *gin.Context)
	DownloadFile(c *gin.Context)
	Health(c *gin.Context)
}

func SetupRouter(h Handler) *gin.Engine {
	router := gin.Default()

	// Навешиваем реальные методы хендлера на маршруты
	router.POST("/files/upload", h.UploadFile)
	router.POST("/files/:fileId/convert", h.ConvertFile)
	router.GET("/files/:convertedFileId/download", h.DownloadFile)

	// Health Check
	router.GET("/health", h.Health)

	return router
}
