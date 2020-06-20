package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterApplicationRoutes - routes for app
func RegisterApplicationRoutes(routerEngine *gin.Engine) {

	routerEngine.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			http.StatusOK, "index.tmpl",
			gin.H{
				"title": "Home Page",
			},
		)
	  
	  })
}