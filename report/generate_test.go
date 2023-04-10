package report

import (
	_ "embed"
	"testing"
)

//go:embed template/report_test.txt
var report []byte

func TestGenerateReport(t *testing.T) {
	result, err := GenerateReport(string(report))
	if err != nil || len(result) == 0 {
		t.Errorf("generation of report failed: %v\n", err.Error())
	}
}
