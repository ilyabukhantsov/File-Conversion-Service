package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func FileUpload(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
	log.Printf("It works")
}
func ConvertFileById(c *gin.Context) {
	id := c.Param("fileId")
	c.String(http.StatusOK, "Hello "+id)
	log.Printf("It works")
}
func GetInfoConvertFileById(c *gin.Context) {
	id := c.Param("convertedFileId")
	c.String(http.StatusOK, "Hello "+id)
	log.Printf("It works")
}
func DownloadFileById(c *gin.Context) {
	id := c.Param("convertedFileId")
	c.String(http.StatusOK, "Hello "+id)
	log.Printf("It works")
}
