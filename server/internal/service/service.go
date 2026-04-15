package service

type Conventor interface {
	HealthCheck() string
	Convert(path string) (status string, newFilePath string)
	FileCheck(path string) error
}

type Service struct {
	Conventor
}

func NewService(c Conventor) *Service {
	return &Service{Conventor: c}
}
