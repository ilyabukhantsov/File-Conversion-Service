package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	ErrNotFound    = errors.New("file not found")
	ErrInvalidID   = errors.New("invalid file id")
	ErrUnsupported = errors.New("unsupported file type or format")
)

var inputExts = map[string]bool{
	".doc": true, ".docx": true, ".odt": true, ".rtf": true, ".txt": true,
	".xls": true, ".xlsx": true, ".ods": true,
	".ppt": true, ".pptx": true, ".odp": true,
}

var outputFormats = map[string]bool{
	"pdf": true, "docx": true, "odt": true, "rtf": true, "txt": true, "html": true,
}

var idPattern = regexp.MustCompile(`^[0-9a-f]{32}$`)

type FileConverter interface {
	Convert(inputPath, outDir, format string) (string, error)
}

type Service struct {
	converter FileConverter
	dataDir   string
}

func NewService(c FileConverter, dataDir string) (*Service, error) {
	abs, err := filepath.Abs(dataDir)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(abs, 0o755); err != nil {
		return nil, err
	}
	return &Service{converter: c, dataDir: abs}, nil
}

func newID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func (s *Service) fileByID(id string) (string, error) {
	if !idPattern.MatchString(id) {
		return "", ErrInvalidID
	}
	dir := filepath.Join(s.dataDir, id)
	entries, err := os.ReadDir(dir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return "", ErrNotFound
		}
		return "", err
	}
	for _, e := range entries {
		if !e.IsDir() {
			return filepath.Join(dir, e.Name()), nil
		}
	}
	return "", ErrNotFound
}

func (s *Service) Upload(name string, src io.Reader) (string, error) {
	name = filepath.Base(strings.ReplaceAll(name, "\\", "/"))
	if name == "." || name == "/" || name == "" {
		return "", ErrUnsupported
	}
	if !inputExts[strings.ToLower(filepath.Ext(name))] {
		return "", ErrUnsupported
	}

	id, err := newID()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(s.dataDir, id)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}

	dst, err := os.Create(filepath.Join(dir, name))
	if err != nil {
		os.RemoveAll(dir)
		return "", err
	}
	if _, err := io.Copy(dst, src); err != nil {
		dst.Close()
		os.RemoveAll(dir)
		return "", err
	}
	if err := dst.Close(); err != nil {
		os.RemoveAll(dir)
		return "", err
	}
	return id, nil
}

func (s *Service) Convert(id, format string) (string, error) {
	format = strings.ToLower(format)
	if !outputFormats[format] {
		return "", ErrUnsupported
	}

	in, err := s.fileByID(id)
	if err != nil {
		return "", err
	}

	newFileID, err := newID()
	if err != nil {
		return "", err
	}
	outDir := filepath.Join(s.dataDir, newFileID)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return "", err
	}

	if _, err := s.converter.Convert(in, outDir, format); err != nil {
		os.RemoveAll(outDir)
		return "", err
	}
	return newFileID, nil
}

func (s *Service) Download(id string) (string, error) {
	return s.fileByID(id)
}
