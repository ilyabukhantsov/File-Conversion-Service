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
	dir := "./files"
	path := dir + "/test.docx"
	newFilePath := "./test.pdf"
	if err := os.Mkdir(dir, 0755); err != nil {
		log.Println("Critical Problem on creating the dir! ", err)
		return "FAIL"
	}
	defer os.RemoveAll("./files")
	f, err := os.Create(path)
	if err != nil {
		return "FAIL"
	}
	f.Close()
	cmd := exec.Command("soffice", "--headless", "--convert-to", "pdf", path)
	if err := cmd.Run(); err != nil {
		log.Println("Critical Problem on executing the conventor! ", err)
		return "FAIL"
	}
	//Deleting old file
	if err := os.Remove(path); err != nil {
		log.Println("Critical Problem on removing the old file! ", err)
		return "FAIL"
		//Deleting new file
	}
	if err := os.Remove(newFilePath); err != nil {
		log.Println("Critical Problem on deleting new file!  ", err)
		return "FAIL"
	}
	fmt.Println("LibreOffice PASS")
	return string("PASS")
}

// TODO Realise it
func (c LibreOffice) Convert(path string) (status string, newFilePath string) {
	fmt.Println("Converted")
	return "200", "file/file.dox"
}
