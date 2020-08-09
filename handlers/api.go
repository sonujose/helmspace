package routers

import (
	"github.com/gin-gonic/gin"
)

//RegisterAPIEndpoints : Register routes for all api endpoints
func RegisterAPIEndpoints(routerEngine *gin.Engine) {
	
	apiHandler := handler{}

	// Setup route group for the API
	api := routerEngine.Group("/api/v1")
	{
		api.GET("/readme/:name/:version", apiHandler.GetChartDetailReadme)
	}

}
