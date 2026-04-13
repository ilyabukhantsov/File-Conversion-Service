package libreoffice

import "fmt"

type LibreOffice struct{}

// TODO Realise it
func (c LibreOffice) HealthCheck() string {
	fmt.Println("All Good")
	return string("all good")
}

// TODO Realise it
func (c LibreOffice) Convert(path string) (status string, newFilePath string) {
	fmt.Println("Converted")
	return "200", "file/file.dox"
}
