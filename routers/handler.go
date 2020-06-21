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
	charts := controller.GetCharts()

	data := gin.H{"title": "Helm-Dimensions", "chartData": charts}

	render(c, data, "index.tmpl")
}

func (h *handler) ShowChartPage(c *gin.Context) {

	var chartData models.ChartItem

	err := c.ShouldBindUri(&chartData); 

	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	fmt.Printf("Chart name : %s \n" + chartData.Name)
	chartItem := controller.GetChartMetadata(chartData.Name)

	data := gin.H{"title": "Helm-Dimensions", "chartItem": chartItem}

	render(c, data, "chart.tmpl")
}

func render(c *gin.Context, data gin.H, templateName string) {
	c.HTML(http.StatusOK, templateName, data)
}