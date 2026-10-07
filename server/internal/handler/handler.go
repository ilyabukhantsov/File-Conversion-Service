package handler

import (
	"errors"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"GoGameV3/internal/dto"
	"GoGameV3/internal/service"
)

const maxUploadSize = 50 << 20 // 50 MB

type Service interface {
	Upload(name string, src io.Reader) (string, error)
	Convert(id, format string) (string, error)
	Download(id string) (string, error)
}

type Handler struct {
	svc Service
}

func NewHandler(s Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) UploadFile(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)

	fh, err := c.FormFile("file")
	if err != nil {
		var tooBig *http.MaxBytesError
		if errors.As(err, &tooBig) {
			c.JSON(http.StatusRequestEntityTooLarge, dto.ErrorResponse{Error: "file too large"})
			return
		}
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: `form field "file" is required`})
		return
	}

	f, err := fh.Open()
	if err != nil {
		writeError(c, err)
		return
	}
	defer f.Close()

	id, err := h.svc.Upload(fh.Filename, f)
	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusCreated, dto.UploadResponse{FileID: id, Filename: fh.Filename})
}

func (h *Handler) ConvertFile(c *gin.Context) {
	format := c.DefaultQuery("format", "pdf")

	newID, err := h.svc.Convert(c.Param("id"), format)
	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.ConvertResponse{ConvertedFileID: newID})
}

func (h *Handler) DownloadFile(c *gin.Context) {
	path, err := h.svc.Download(c.Param("id"))
	if err != nil {
		writeError(c, err)
		return
	}

	c.FileAttachment(path, filepath.Base(path))
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func writeError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrInvalidID):
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid file id"})
	case errors.Is(err, service.ErrUnsupported):
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "unsupported file type or format"})
	case errors.Is(err, service.ErrNotFound):
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "file not found"})
	default:
		log.Printf("internal error: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal error"})
	}
}
