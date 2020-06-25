package routers

import (

	"github.com/gin-gonic/gin"
)

// InitilaizeRoutes - routes for the application
func InitilaizeRoutes(router *gin.Engine) {
	
	apiHandler := handler{}

	router.GET("/", apiHandler.ShowIndexPage)

	router.GET("/chart/:name", apiHandler.ShowChartPage)
}