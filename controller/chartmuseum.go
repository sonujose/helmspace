package controller

import (
	"net/http"
	"log"
	"io/ioutil"

	"github.com/helm-dimensions/models"
)

// DEFAULT: API Endpoint for chartmuseum server
var api = "/api/charts"

// GetCharts - Fetch Charts from Chartmuseum server
func GetCharts(repoEndpoint string) map[string][]models.Chart {

	repoAPI := repoEndpoint + api

	log.Printf("Fetching the url %v \n", repoAPI)

	response,err := http.Get(repoAPI)

	if err != nil {
		log.Fatalf("Unable to retrieve chart data from %v - Error: %v", repoAPI, err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	charts, err := models.GetNewCharts(data)

	if err != nil {
		log.Printf("Unable to marshell data %v", err)
	}

	return charts
}

//GetChartMetadata - Get Metadata from Chartmuseum server
func GetChartMetadata(chartName string, repoEndpoint string) []models.Chart {
	
	repoDetailAPI := repoEndpoint + api + "/" + chartName

	log.Printf("Fetching the chartMetadata from  %v \n", repoDetailAPI)

	response,err := http.Get(repoDetailAPI)

	if err != nil {
		log.Fatalf("Unable to retrieve chart data from %v - Error: %v", repoDetailAPI, err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	chartItem, err := models.GetNewChartItem(data)

	if err != nil {
		log.Printf("Unable to marshell data %v", err)
	}

	return chartItem

}
