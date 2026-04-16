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
		"fileId": "mock200",
	})
}
func MockConvertFileById(c *gin.Context) {
	switch id := c.Param("fileId"); id {
	case "mock200":
		c.JSON(200, map[string]any{
			"convertedFileId": "def456",
			"status":          "success",
		})
	case "mock400":
		c.JSON(400, map[string]any{
			"code":    "BAD_REQUEST",
			"message": "Invalid input",
			"details": map[string]any{
				"additionalProp1": map[string]any{},
			},
		})
	case "mock404":
		c.JSON(404, map[string]any{
			"code":    "BAD_REQUEST",
			"message": "Invalid input",
			"details": map[string]any{
				"additionalProp1": map[string]any{},
			},
		})
	case "mock422":
		c.JSON(422, map[string]any{
			"code":    "BAD_REQUEST",
			"message": "Invalid input",
			"details": map[string]any{
				"additionalProp1": map[string]any{},
			},
		})
	case "mock500":
		c.JSON(500, map[string]any{
			"code":    "BAD_REQUEST",
			"message": "Invalid input",
			"details": map[string]any{
				"additionalProp1": map[string]any{},
			},
		})
	default:
		c.JSON(456, map[string]any{
			"code":    "BAD_REQUEST",
			"message": "Invalid input",
			"details": "No validation for other id's in MOCK server, check DOCS for see correct id's for tests!",
		})
	}
}

func MockDownloadFileById(c *gin.Context) {
	switch id := c.Param("convertedFileId"); id {
	case "mock200":
		c.JSON(200, gin.H{
			"content": "no file, sry, its mock server for a reason!",
		})
	case "mock404":
		c.JSON(404, map[string]any{
			"code":    "BAD_REQUEST",
			"message": "Invalid input",
			"details": map[string]any{
				"additionalProp1": map[string]any{},
			},
		})
	case "mock500":
		c.JSON(500, map[string]any{
			"code":    "BAD_REQUEST",
			"message": "Invalid input",
			"details": map[string]any{
				"additionalProp1": map[string]any{},
			},
		})
	default:
		c.JSON(456, map[string]any{
			"code":    "BAD_REQUEST",
			"message": "Invalid input",
			"details": "No validation for other id's in MOCK server, check DOCS for see correct id's for tests!",
		})
	}
}
