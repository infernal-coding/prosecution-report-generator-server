package main

import (
	"github.com/infernal-coding/prosecution-report-generator-server/web"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/api/v1/status/health", web.HandleHealth)

	log.Fatal(http.ListenAndServe(":8080", router))
}
