package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterStatusEndpoints(group *gin.RouterGroup) {
	v1 := group.Group("/v1/status")
	v1.GET("/health", getStatusHealth)
}

func getStatusHealth(c *gin.Context) {
	c.String(http.StatusOK, "UP")
}
