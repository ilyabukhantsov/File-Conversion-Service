package libreoffice

type LibreOffice struct{}

func NewConverter() *LibreOffice {
	return &LibreOffice{}
}

func (c LibreOffice) Convert(inputPath string) (string, error) {
	return " ", nil
}
