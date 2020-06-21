package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helm-dimensions/controller"
)

type handler struct {}


func (h *handler) ShowIndexPage(c *gin.Context) {
	charts := controller.GetCharts()

	data := gin.H{"title": "Helm-Dimensions", "chartData": charts}

	render(c, data, "index.tmpl")
}

func (h *handler) ShowChartPage(c *gin.Context) {
	charts := controller.GetCharts()

	data := gin.H{"title": "Helm-Dimensions", "chartData": charts}

	render(c, data, "index.tmpl")
}

func render(c *gin.Context, data gin.H, templateName string) {
	c.HTML(http.StatusOK, templateName, data)
}