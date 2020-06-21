package controller

import (
	"net/http"
	"log"
	"io/ioutil"
	"os"
	"fmt"

	"github.com/helm-dimensions/models"
)

// API Endpoint for chartmuseum server
var api = "/api/charts"

// GetCharts - Fetch sm charts
func GetCharts() map[string][]models.Chart {

	urlfull := getBaseURL()

	log.Printf("Fetching the url %v \n", urlfull)

	response,err := http.Get(getBaseURL())

	if err != nil {
		log.Fatalf("Unable to retrieve chart data from %v - Error: %v", getBaseURL(), err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	charts, err := models.GetNewCharts(data)

	if err != nil {
		log.Printf("Unable to marshell data %v", err)
	}

	return charts
}

//GetChartMetadata - Get Metadata of a particular chart
func GetChartMetadata(chartName string) []models.Chart {
	urlfull := getBaseURL() + "/" + chartName

	log.Printf("Fetching the chartMetadata from  %v \n", urlfull)

	response,err := http.Get(urlfull)

	if err != nil {
		log.Fatalf("Unable to retrieve chart data from %v - Error: %v", urlfull, err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))

	chartItem, err := models.GetNewChartItem(data)

	if err != nil {
		log.Printf("Unable to marshell data %v", err)
	}

	return chartItem

}

func getBaseURL() string {

	baseURL := "http://pickles-charts.australiaeast.cloudapp.azure.com"

	apiendpoint := os.Getenv("CHART_MUSEUM_URL")

	if apiendpoint != "" {
		log.Printf("Chartmuseum server provided - %s", apiendpoint)
		baseURL = apiendpoint
	}

	url := baseURL + api

	return url

}
