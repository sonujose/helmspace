package controller

import (
	"net/http"
	"log"
	"io/ioutil"
	"os"
	"fmt"

	"github.com/chartsmirror/models"
)

var api = "/api/charts"

// GetCharts - Fetch sm charts
func GetCharts() map[string][]models.Chart {

	urlfull := getBaseURL()

	fmt.Printf("Fetching the url %v", urlfull)

	response,err := http.Get(getBaseURL())

	if err != nil {
		log.Fatalf("Unable to retrieve chart data from %v - Error: %v", getBaseURL(), err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	//fmt.Println(string(data))

	charts, err := models.GetNewCharts(data)

	if err != nil {
		log.Printf("Unable to marshell data %v", err)
	}

	return charts
}

func getBaseURL() string {

	baseURL := "http://pickles-charts.australiaeast.cloudapp.azure.com"

	apiendpoint := os.Getenv("CHART_MUSEUM_API_GET_CHARTS")

	if apiendpoint != "" {
		baseURL = apiendpoint
	}

	url := baseURL + api

	return url

}

