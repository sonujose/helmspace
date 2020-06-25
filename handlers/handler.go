package routers

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/helm-dimensions/controller"
	"github.com/helm-dimensions/models"
)

type handler struct {}


func (h *handler) ShowIndexPage(c *gin.Context) {

	repoItem := controller.GetRepoDetails()

	charts := controller.GetCharts(repoItem.URL)

	data := gin.H{"title": "Helm-Dimensions", "chartData": charts}

	render(c, data, "index.tpl")
}

func (h *handler) ShowChartPage(c *gin.Context) {

	var chartData models.ChartItem

	err := c.ShouldBindUri(&chartData); 

	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	fmt.Printf("Chart name selected - %s \n", chartData.Name)

	repoItem := controller.GetRepoDetails()
	
	chartItem := controller.GetChartMetadata(chartData.Name, repoItem.URL)

	data := gin.H{"title": "Helm-Dimensions", "chartItem": chartItem, "repoDetails": repoItem}

	render(c, data, "chart.tpl")
}

func render(c *gin.Context, data gin.H, templateName string) {
	c.HTML(http.StatusOK, templateName, data)
}