package libreoffice

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestLibreOfficeConvert(t *testing.T) {
	outDir := t.TempDir()

	inputDir := t.TempDir()
	inputPath := filepath.Join(inputDir, "test.txt")

	if err := os.WriteFile(inputPath, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}

	binary := createFakeSoffice(t, `
		outDir := ""
		input := ""

		for i := 0; i < len(os.Args); i++ {
			if os.Args[i] == "--outdir" && i+1 < len(os.Args) {
				outDir = os.Args[i+1]
			}

			if i == len(os.Args)-1 {
				input = os.Args[i]
			}
		}

		base := strings.TrimSuffix(
			filepath.Base(input),
			filepath.Ext(input),
		)

		result := filepath.Join(outDir, base+".pdf")

		if err := os.WriteFile(result, []byte("converted"), 0644); err != nil {
			os.Exit(1)
		}
	`)

	converter := &LibreOffice{
		binary:  binary,
		timeout: time.Second,
		sem:     make(chan struct{}, maxParallelConvert),
	}

	result, err := converter.Convert(inputPath, outDir, "pdf")
	if err != nil {
		t.Fatalf("Convert() error = %v", err)
	}

	expected := filepath.Join(outDir, "test.pdf")

	if result != expected {
		t.Errorf("Convert() = %q, want %q", result, expected)
	}

	data, err := os.ReadFile(result)
	if err != nil {
		t.Fatalf("failed to read result: %v", err)
	}

	if string(data) != "converted" {
		t.Errorf("result content = %q, want %q", string(data), "converted")
	}
}

func TestLibreOfficeConvertSofficeError(t *testing.T) {
	binary := createFakeSoffice(t, `
		os.Exit(1)
	`)

	converter := &LibreOffice{
		binary:  binary,
		timeout: time.Second,
		sem:     make(chan struct{}, maxParallelConvert),
	}

	outDir := t.TempDir()
	inputPath := filepath.Join(t.TempDir(), "test.txt")

	if err := os.WriteFile(inputPath, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}

	_, err := converter.Convert(inputPath, outDir, "pdf")

	if err == nil {
		t.Fatal("Convert() expected error, got nil")
	}

	if !strings.Contains(err.Error(), "soffice failed") {
		t.Errorf("error = %q, want 'soffice failed'", err)
	}
}

func TestLibreOfficeConvertNoOutput(t *testing.T) {
	binary := createFakeSoffice(t, `
		// Do nothing.
	`)

	converter := &LibreOffice{
		binary:  binary,
		timeout: time.Second,
		sem:     make(chan struct{}, maxParallelConvert),
	}

	outDir := t.TempDir()
	inputPath := filepath.Join(t.TempDir(), "test.txt")

	if err := os.WriteFile(inputPath, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}

	_, err := converter.Convert(inputPath, outDir, "pdf")

	if err == nil {
		t.Fatal("Convert() expected error, got nil")
	}

	if !strings.Contains(err.Error(), "produced no output file") {
		t.Errorf(
			"error = %q, want 'produced no output file'",
			err,
		)
	}
}

func TestLibreOfficeConvertTimeout(t *testing.T) {
	binary := createFakeSoffice(t, `
		time.Sleep(5 * time.Second)
	`)

	converter := &LibreOffice{
		binary:  binary,
		timeout: 50 * time.Millisecond,
		sem:     make(chan struct{}, maxParallelConvert),
	}

	outDir := t.TempDir()
	inputPath := filepath.Join(t.TempDir(), "test.txt")

	if err := os.WriteFile(inputPath, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}

	start := time.Now()

	_, err := converter.Convert(inputPath, outDir, "pdf")

	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("Convert() expected timeout error, got nil")
	}

	if elapsed > time.Second {
		t.Errorf("Convert() took too long: %v", elapsed)
	}

	if !strings.Contains(err.Error(), "soffice failed") {
		t.Errorf("error = %q, want 'soffice failed'", err)
	}
}

func TestLibreOfficeConvertOutputName(t *testing.T) {
	binary := createFakeSoffice(t, `
		outDir := ""

		for i := 0; i < len(os.Args); i++ {
			if os.Args[i] == "--outdir" && i+1 < len(os.Args) {
				outDir = os.Args[i+1]
			}
		}

		input := os.Args[len(os.Args)-1]

		base := strings.TrimSuffix(
			filepath.Base(input),
			filepath.Ext(input),
		)

		result := filepath.Join(outDir, base+".docx")

		if err := os.WriteFile(result, []byte("test"), 0644); err != nil {
			os.Exit(1)
		}
	`)

	converter := &LibreOffice{
		binary:  binary,
		timeout: time.Second,
		sem:     make(chan struct{}, maxParallelConvert),
	}

	outDir := t.TempDir()
	inputPath := filepath.Join(t.TempDir(), "my-document.txt")

	if err := os.WriteFile(inputPath, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}

	result, err := converter.Convert(inputPath, outDir, "docx")
	if err != nil {
		t.Fatalf("Convert() error = %v", err)
	}

	expected := filepath.Join(outDir, "my-document.docx")

	if result != expected {
		t.Errorf("result = %q, want %q", result, expected)
	}
}

func createFakeSoffice(t *testing.T, code string) string {
	t.Helper()

	dir := t.TempDir()
	file := filepath.Join(dir, "soffice")

	content := `#!/bin/sh
exec go run "` + filepath.Join(dir, "fake.go") + `"
`

	if err := os.WriteFile(file, []byte(content), 0o755); err != nil {
		t.Fatal(err)
	}

	fakeGo := `
package main

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	` + code + `
}
`

	if err := os.WriteFile(
		filepath.Join(dir, "fake.go"),
		[]byte(fakeGo),
		0o644,
	); err != nil {
		t.Fatal(err)
	}

	return file
}
