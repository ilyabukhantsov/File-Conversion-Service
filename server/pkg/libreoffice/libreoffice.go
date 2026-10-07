// Package libreoffice converts documents by running LibreOffice in headless mode.
package libreoffice

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	defaultTimeout     = 2 * time.Minute
	maxParallelConvert = 2
)

type LibreOffice struct {
	binary  string
	timeout time.Duration
	sem     chan struct{}
}

func NewConverter() *LibreOffice {
	return &LibreOffice{
		binary:  "soffice",
		timeout: defaultTimeout,
		sem:     make(chan struct{}, maxParallelConvert),
	}
}

func (c *LibreOffice) Convert(inputPath, outDir, format string) (string, error) {
	c.sem <- struct{}{}
	defer func() { <-c.sem }()

	profile, err := os.MkdirTemp("", "lo-profile-*")
	if err != nil {
		return "", fmt.Errorf("create profile dir: %w", err)
	}
	defer os.RemoveAll(profile)

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, c.binary,
		"--headless",
		"--norestore",
		"--nolockcheck",
		"--nodefault",
		"--nofirststartwizard",
		"-env:UserInstallation=file://"+profile,
		"--convert-to", format,
		"--outdir", outDir,
		inputPath,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("soffice failed: %w: %s", err, strings.TrimSpace(string(out)))
	}

	base := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	result := filepath.Join(outDir, base+"."+format)
	if _, err := os.Stat(result); err != nil {
		return "", fmt.Errorf("soffice produced no output file: %w", err)
	}
	return result, nil
}
