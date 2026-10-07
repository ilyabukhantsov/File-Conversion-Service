package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"GoGameV3/internal/service"
)

type Service interface {
	Convert(path string, s Service) (newFilePath string, err error)
	Upload(s Service) (newFilePath string, err error)
	Download(path string, s Service) (err error)
}

type HandlerImpl struct {
	service *service.Service
}

func NewHandler(s *service.Service) *HandlerImpl {
	return &HandlerImpl{
		service: s,
	}
}

func (h *HandlerImpl) Health(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *HandlerImpl) UploadFile(c *gin.Context) {
	// h.service.Upload(...)
}

func (h *HandlerImpl) ConvertFile(c *gin.Context) {
	// h.service.Convert(...)
}

func (h *HandlerImpl) DownloadFile(c *gin.Context) {
	// h.service.Download(...)
}
