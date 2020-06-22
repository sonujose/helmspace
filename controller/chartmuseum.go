package controller

import (
	"net/http"
	"log"
	"io/ioutil"
	"os"

	"github.com/helm-dimensions/models"
)

// DEFAULT: API Endpoint for chartmuseum server
var api = "/api/charts"
var repoURL = "http://localhost:9000"

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

//GetRepoDetails - details and name of the chart Repo
//TODO: Current Implemtation is limited to only one repo server 
// - Future need to read from a configuration file based on chart data
func GetRepoDetails() models.Repo {

	RepoCM := os.Getenv("CHART_MUSEUM_URL")
	
	if RepoCM != "" {
		log.Printf("Fetching repo from CHART_MUSEUM_URL - %s", RepoCM)
		repoURL = RepoCM
	}

	repo := models.Repo{ Name: "demo", URL: repoURL }
	log.Printf("Repository - %s", repo.URL)

	return repo
}

