package controller

import(
	"log"

	"github.com/sonujose/helmspace/utilities"
	"github.com/sonujose/helmspace/models"
)

// FetchChartReadme - Get Chartdate readme content
func FetchChartReadme(repoURL string, chartData models.ChartData) *string {

	log.Printf("Creating tarball url -  %s, %s, %s ", repoURL, chartData.Name, chartData.Version)

	tarballURL := getChartTarballURL(repoURL, chartData.Name, chartData.Version)
	
	log.Printf("Fetching chart data from %s & README from %s/README.md", tarballURL, chartData.Name)

	readmeContent, err := utilities.GetFileBlobFromTarBall(tarballURL, utilities.CreateKeyString(chartData.Name ,"/README.md"))

	if err != nil {
		log.Fatalf("Unable to retrieve chart data from %v - Error: %v", tarballURL, err)
	}

	return readmeContent
}

func getChartTarballURL(repoURL string, chartName string, version string) string {
	
	return utilities.CreateKeyString(repoURL ,"/charts/" ,chartName ,"-" ,version ,".tgz")

}