package web

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestRegisterStatusEndpoints(t *testing.T) {
	engine := gin.New()
	routerGroup := engine.Group("/api")

	RegisterStatusEndpoints(routerGroup)

	testIfRoutePresent(t, engine, "GET", "/api/v1/status/health")
}

func testIfRoutePresent(t *testing.T, engine *gin.Engine, method string, path string) {
	for _, route := range engine.Routes() {
		if route.Method == method && route.Path == path {
			return
		}
	}
	t.Errorf("Route '[%v] %v' is not registered on engine!", method, path)
}

func TestGetStatusHealth(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseRecorder)

	GetStatusHealth(context)

	resultCode := responseRecorder.Code
	resultBody := responseRecorder.Body.String()
	if resultCode != 200 || resultBody != "UP" {
		t.Errorf("Expected response '200 - UP' but got '%v - %v'", resultCode, resultBody)
	}
}
