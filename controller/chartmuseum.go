package controller

import (
	"net/http"
	"log"
	"io/ioutil"
	"os"

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

func getBaseURL() string {

	baseURL := "http://localhost:9000"

	apiendpoint := os.Getenv("CHART_MUSEUM_URL")

	if apiendpoint != "" {
		log.Printf("Chartmuseum server provided - %s", apiendpoint)
		baseURL = apiendpoint
	}

	url := baseURL + api

	return url

}

