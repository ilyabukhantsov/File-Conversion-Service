package libreoffice

import (
	"os"
	"testing"
)

func TestIsExist(t *testing.T) {
	result := isExist("../../files/mock.docx")
	var expected error = nil

	if result != expected {
		t.Errorf("Очікували %d, але отримали %d", expected, result)
	}
}

func TestConvert(t *testing.T) {
	newFilePath, err := LibreOffice{}.Convert("../../files/mock.docx", "../../files/mock.pdf")

	var expectedError error = nil
	expectedNewFilePath := "../../files/mock.pdf"

	if err != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, err)
	}

	if newFilePath != expectedNewFilePath {
		t.Fatalf("expected %v, got %v", expectedNewFilePath, newFilePath)
	}

	if _, err := os.Stat(newFilePath); err != nil {
		t.Fatalf("pdf file was not created: %v", err)
	}

	if err := os.RemoveAll(newFilePath); err != nil {
		t.Fatalf("expected %v, got %v", expectedError, err)
	}
}
