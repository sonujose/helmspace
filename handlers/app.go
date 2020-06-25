package routers

import (

	"github.com/gin-gonic/gin"
)

// ReisterApplicationRoutes - routes for the application
func ReisterApplicationRoutes(router *gin.Engine) {
	
	apiHandler := handler{}

	router.GET("/", apiHandler.ShowIndexPage)

	router.GET("/chart/:name", apiHandler.ShowChartPage)
}