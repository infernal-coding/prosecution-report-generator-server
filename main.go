package main

import (
	"github.com/gin-gonic/gin"
	"github.com/infernal-coding/prosecution-report-generator-server/report"
	"github.com/infernal-coding/prosecution-report-generator-server/web"
	"log"
)

func main() {
	report.SetupTemplate()

	router := gin.Default()

	api := router.Group("/api")
	web.RegisterStatusEndpoints(api)
	web.RegisterReportEndpoints(api)

	log.Fatal(router.Run())
}
