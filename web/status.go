package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStatusHealth(c *gin.Context) {
	c.String(http.StatusOK, "UP")
}
