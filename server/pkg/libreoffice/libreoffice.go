package libreoffice

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type LibreOffice struct{}

// FileCheck implements [service.Conventor].
func (c LibreOffice) FileCheck(path string) error {
	panic("unimplemented")
}

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

func isExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c LibreOffice) Convert(path string, outputPath string) (newFilePath string, err error) {
	if err := isExist(path); err != nil {
		return "", err
	}

	cmd := exec.Command("soffice", "--headless", "--convert-to", "pdf", "--outdir", outputPath, path)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	newFilePath = strings.TrimSuffix(path, filepath.Ext(path)) + ".pdf"
	return newFilePath, nil
}
