package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type MockHandler struct{}

func MockFileUpload(c *gin.Context) {
	file, err := c.FormFile("upload")
	if err != nil {
		fmt.Println()
		c.JSON(500, map[string]any{
			"code":    "BAD_REQUEST",
			"message": "Invalid input",
			"details": map[string]any{
				"additionalProp1": map[string]any{},
			},
		})
		return
	}

	_ = file

	c.JSON(200, gin.H{
		"fileId": "abc123",
	})
}
func MockConvertFileById(c *gin.Context) {

}

func MockDownloadFileById(c *gin.Context) {
	//TODO implimication of download mock
}
