package libreoffice

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type LibreOffice struct{}

// TODO Make humanlike realisation
func (c LibreOffice) HealthCheck() string {
	file, err := os.Create("test.docx")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("soffice", "--headless", "--convert-to", "pdf", file.Name())
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	if err := os.Remove(file.Name()); err != nil {
		log.Fatalln(err)
	}
	if err := os.Remove("test.pdf"); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("LibreOffice PASS")
	return string("PASS")
}

// TODO Realise it
func (c LibreOffice) Convert(path string) (status string, newFilePath string) {
	fmt.Println("Converted")
	return "200", "file/file.dox"
}
