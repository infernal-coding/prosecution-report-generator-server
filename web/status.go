package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterStatusEndpoints registers all endpoints for the status group on the given router group.
func RegisterStatusEndpoints(group *gin.RouterGroup) {
	v1 := group.Group("/v1/status")
	v1.GET("/health", GetStatusHealth)
}

// GetStatusHealth provides an endpoint that responds with a status of 200 as the server is successfully running.
func GetStatusHealth(c *gin.Context) {
	c.String(http.StatusOK, "UP")
}
