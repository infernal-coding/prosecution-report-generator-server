package report

import (
	"os"
	"os/exec"
	"path"
)

const (
	DirName = "report"
	TexName = "report.tex"
	PdfName = "report.pdf"
)

// GenerateReport generates the pdf file for the given report with pdflatex.
func GenerateReport(report string) ([]byte, error) {
	tmpDir, err := createReportTexFile(report)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmpDir)

	return buildReportFile(tmpDir)
}

func createReportTexFile(report string) (string, error) {
	dir, err := os.MkdirTemp("", DirName)
	if err != nil {
		return "", err
	}

	filePath := path.Join(dir, TexName)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		os.RemoveAll(dir)
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString(report)
	if err != nil {
		os.RemoveAll(dir)
		return "", err
	}

	return dir, err
}

func buildReportFile(dir string) ([]byte, error) {
	for i := 0; i < 2; i++ {
		cmd := exec.Command("pdflatex", "-interaction=nonstopmode", "-halt-on-error", TexName)
		cmd.Dir = dir

		err := cmd.Run()
		if err != nil {
			return nil, err
		}
	}

	reportFile := path.Join(dir, PdfName)
	generatedReport, err := os.ReadFile(reportFile)
	if err != nil {
		return nil, err
	}

	return generatedReport, nil
}
