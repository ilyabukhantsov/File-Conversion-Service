package router

import (
	"GoGameV3/internal/handler"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/files/upload", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
		log.Printf("It works")
	})
	router.POST("/files/:fileId/convert", func(c *gin.Context) {
		id := c.Param("fileId")
		c.String(http.StatusOK, "Hello "+id)
		log.Printf("It works")
	})
	router.GET("/files/:convertedFileId", func(c *gin.Context) {
		id := c.Param("convertedFileId")
		c.String(http.StatusOK, "Hello "+id)
		log.Printf("It works")
	})
	router.GET("/files/:convertedFileId/download", func(c *gin.Context) {
		id := c.Param("convertedFileId")
		c.String(http.StatusOK, "Hello "+id)
		log.Printf("It works")
	})

	return router
}
