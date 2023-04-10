package report

import (
	"bytes"
	_ "embed"
	"github.com/infernal-coding/prosecution-report-generator-server/dto"
	"testing"
)

//go:embed template/report_test.txt
var expectedResult []byte

func TestSetupTemplate(t *testing.T) {
	SetupTemplate()

	if reportTemplate == nil {
		t.Error("report template was not initialized")
		return
	}
}

func TestExecuteTemplate(t *testing.T) {
	var result bytes.Buffer
	data := dto.ReportData{
		Header: dto.ReportHeader{
			Team:            "Herren",
			Home:            "SC Haudaneben",
			Away:            "FC Krumme Haxn",
			Date:            "01.07.2022",
			Place:           "Grossköllnbach",
			HomeScoreHalf:   0,
			AwayScoreHalf:   1,
			HomeScore:       2,
			AwayScore:       1,
			MatchIdentifier: "12345",
			MatchClass: dto.MatchClass{
				Private: true,
			},
		},
		Content: dto.ReportContent{
			Subject:    "Fehlendes Passfoto",
			PassportId: "42",
			Who:        "Max Mustermann (8) vom SC Haudaneben",
			Other:      "Bei Herrn Mustermann fehlte bis nach Spielende das Foto im ESB.",
		},
		Footer: dto.ReportFooter{
			Name:   "Franz Huber",
			Street: "Hauptstraße 1",
			Zip:    "94315",
			Place:  "Straubing",
			Group:  "Straubing",
			Mail:   "huber@test.de",
			Recipients: dto.ReportRecipients{
				GSO: true,
			},
		},
	}

	err := ExecuteTemplate(&result, &data)

	if err != nil {
		t.Errorf("template execution should succeed but failed with error: %v", err.Error())
		return
	}

	if result.String() != string(expectedResult) {
		t.Errorf("resulting report is not matching expecting result, generated report:\n%v", result.String())
	}
}

func TestExecuteTemplate_NotInitialized(t *testing.T) {
	reportTemplate = nil

	var result bytes.Buffer
	data := dto.ReportData{}

	err := ExecuteTemplate(&result, &data)

	if err == nil {
		t.Error("executing with uninitialized template should not work")
	}
}
