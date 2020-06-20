package routers

import (
	"net/http"
	//"encoding/json"
	//"log"

	"github.com/gin-gonic/gin"
	"github.com/chartsmirror/controller"

)

// InitilaizeRoutes - routes for app
func InitilaizeRoutes(router *gin.Engine) {
	
	router.GET("/", showIndexPage)
}

func showIndexPage(c *gin.Context) {

	charts := controller.GetCharts()

	// jsonString, err := json.Marshal(charts)
	
	// if err != nil {
	// 	log.Fatalf("Unable to retrive json data from %v", charts)
	// }
	data := gin.H{"title": "Home Page", "chartData": charts}

	render(c, data, "index.tmpl")
}

func render(c *gin.Context, data gin.H, templateName string) {
	c.HTML(http.StatusOK, templateName, data)
}