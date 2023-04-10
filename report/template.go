package report

import (
	_ "embed"
	"errors"
	"github.com/infernal-coding/prosecution-report-generator-server/dto"
	"io"
	"log"
	"text/template"
)

var (
	//go:embed template/report.tex
	templateData   []byte
	reportTemplate *template.Template
)

// SetupTemplate parses the report template and stores it in runtime.
func SetupTemplate() {
	reportTemplate = template.Must(template.New("").Parse(string(templateData)))
	log.Println("Report template parsed successfully!")
}

// ExecuteTemplate executes the report template to the given writer with the given report data.
// If the report was not set up before, this function returns an error.
func ExecuteTemplate(writer io.Writer, data *dto.ReportData) error {
	if reportTemplate == nil {
		return errors.New("template must be setup before executing")
	}
	return reportTemplate.Execute(writer, data)
}
