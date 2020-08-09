package routers

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/sonujose/helmspace/controller"
	"github.com/sonujose/helmspace/models"
)

type handler struct {}


func (h *handler) ShowIndexPage(c *gin.Context) {

	repoItem := controller.GetRepoDetails()

	charts := controller.GetCharts(repoItem.URL)

	data := gin.H{"title": "Helmspace", "chartData": charts}

	render(c, data, "charts/index.tpl")
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

	data := gin.H{"title": "Helmspace", "chartItem": chartItem, "repoDetails": repoItem}

	render(c, data, "charts/chart.tpl")
}

func (h *handler) GetChartDetailReadme(c *gin.Context) {
	var chartData models.ChartData

	err := c.ShouldBindUri(&chartData); 

	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	fmt.Printf("Chart name - %s , version - %s \n", chartData.Name, chartData.Version)

	repoItem := controller.GetRepoDetails()
	
	readmeData := controller.FetchChartReadme(repoItem.URL, chartData)
	
	c.JSON(http.StatusOK, map[string]string{
		"readme": *readmeData,
	})
}

func (h *handler) ShowVisualizePage(c *gin.Context) {

	var chartData models.ChartItem

	err := c.ShouldBindUri(&chartData); 

	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	fmt.Printf("Chart name selected - %s \n", chartData.Name)

	repoItem := controller.GetRepoDetails()
	
	chartItem := controller.GetChartMetadata(chartData.Name, repoItem.URL)

	data := gin.H{"title": "Helmspace", "chartItem": chartItem, "repoDetails": repoItem}

	render(c, data, "visualize/index.tpl")
}


func render(c *gin.Context, data gin.H, templateName string) {
	c.HTML(http.StatusOK, templateName, data)
}