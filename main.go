package main

import (
	"github.com/gin-gonic/gin"
	"github.com/infernal-coding/prosecution-report-generator-server/web"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/api/v1/status/health", web.GetStatusHealth)

	log.Fatal(router.Run())
}
