package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//RegisterAPIEndpoints : Register routes for all api endpoints
func RegisterAPIEndpoints(routerEngine *gin.Engine) {
	
	// Setup route group for the API
	api := routerEngine.Group("/api/v1")
	{
		api.GET("/charts", ListCharts)
	}

}

// ListPods - List all pods in a namespace
func ListCharts(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"pod": "pod1-dssdasd",
	})
}