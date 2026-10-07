package router

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	UploadFile(c *gin.Context)
	ConvertFile(c *gin.Context)
	DownloadFile(c *gin.Context)
	Health(c *gin.Context)
}

func SetupRouter(h Handler, middleware ...gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.Use(middleware...)

	r.POST("/files/upload", h.UploadFile)
	r.POST("/files/:id/convert", h.ConvertFile)
	r.GET("/files/:id/download", h.DownloadFile)

	r.GET("/health", h.Health)

	return r
}
