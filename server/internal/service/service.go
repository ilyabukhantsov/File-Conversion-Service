package service

type FileConverter interface {
	Convert(inputPath string) (string, error)
}

type Service struct {
	converter FileConverter
}

func NewService(c FileConverter) *Service {
	return &Service{
		converter: c,
	}
}

func (s *Service) Convert(path string) (newFilePath string, err error)
func (s *Service) Upload() (newFilePath string, err error)
func (s *Service) Download(path string) (err error)
