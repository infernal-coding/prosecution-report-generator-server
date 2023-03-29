package report

import (
	"errors"
	"github.com/infernal-coding/prosecution-report-generator-server/dto"
	"io"
	"log"
	"text/template"
)

var reportTemplate *template.Template

func SetupTemplate() {
	reportTemplate = template.Must(template.New("").ParseFiles("template/report.tex"))
	log.Println("Report template parsed successfully!")
}

func ExecuteTemplate(writer io.Writer, data *dto.ReportData) error {
	if reportTemplate == nil {
		return errors.New("template must be setup before executing")
	}
	return reportTemplate.ExecuteTemplate(writer, "report.tex", data)
}
