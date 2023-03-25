package main

import (
	"github.com/gin-gonic/gin"
	"github.com/infernal-coding/prosecution-report-generator-server/web"
	"log"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	web.RegisterStatusEndpoints(api)

	log.Fatal(router.Run())
}
