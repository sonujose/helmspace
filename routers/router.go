package routers

import (

	"github.com/gin-gonic/gin"
)

// InitilaizeRoutes - routes for app
func InitilaizeRoutes(router *gin.Engine) {
	
	apiHandler := handler{}

	router.GET("/", apiHandler.ShowIndexPage)

	router.GET("/chart", apiHandler.ShowChartPage)
}