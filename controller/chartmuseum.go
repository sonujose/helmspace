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
var repoURL = "http://localhost:9000"

// GetCharts - Fetch Charts from Chartmuseum server
func GetCharts() map[string][]models.Chart {

	urlfull := getRepoURL()

	log.Printf("Fetching the url %v \n", urlfull)

	response,err := http.Get(urlfull)

	if err != nil {
		log.Fatalf("Unable to retrieve chart data from %v - Error: %v", urlfull, err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	charts, err := models.GetNewCharts(data)

	if err != nil {
		log.Printf("Unable to marshell data %v", err)
	}

	return charts
}

//GetChartMetadata - Get Metadata from Chartmuseum server
func GetChartMetadata(chartName string) []models.Chart {
	urlfull := getRepoURL() + "/" + chartName

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

//TODO: Current Implemtation is limited to only one repo server 
// - Future need to read from a configuration file
func getRepoURL() string {

	customRepoURL := os.Getenv("CHART_MUSEUM_URL")
	
	if customRepoURL != "" {
		log.Printf("Chartmuseum server provided - %s", customRepoURL)
		repoURL = customRepoURL
	}

	url := repoURL + api

	return url
}

