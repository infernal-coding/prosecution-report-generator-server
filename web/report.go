package web

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/infernal-coding/prosecution-report-generator-server/dto"
	"github.com/infernal-coding/prosecution-report-generator-server/report"
	"net/http"
)

func RegisterReportEndpoints(group *gin.RouterGroup) {
	v1 := group.Group("/v1/report")
	v1.POST("/generate", PostReportGenerate)
}

func PostReportGenerate(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var data dto.ReportData
	err = json.Unmarshal(body, &data)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	buffer := new(bytes.Buffer)
	err = report.ExecuteTemplate(buffer, &data)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	generateReport, err := report.GenerateReport(buffer.String())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", generateReport)
}
