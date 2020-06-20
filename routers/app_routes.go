package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitilaizeRoutes - routes for app
func InitilaizeRoutes(router *gin.Engine) {
	
	router.GET("/", showIndexPage)
}

func showIndexPage(c *gin.Context) {

	data := gin.H{"title": "Home Page"}

	render(c, data, "index.tmpl")
}

func render(c *gin.Context, data gin.H, templateName string) {
	c.HTML(http.StatusOK, templateName, data)
}