package controller

import(
	"log"

	"github.com/sonujose/helmer/utilities"
	"github.com/sonujose/helmer/models"
)

// FetchChartReadme - Get Chartdate readme content
func FetchChartReadme(repoURL string, chartVersion models.ChartVersion) *string {

	log.Printf("Creating tarball url -  %s, %s, %s ", repoURL, chartVersion.Name, chartVersion.Version)

	tarballURL := getChartTarballURL(repoURL, chartVersion.Name, chartVersion.Version)
	
	log.Printf("Fetching chart data from %s", tarballURL)

	readmeContent, err := utilities.GetFileBlobFromTarBall(tarballURL, "core-microservice/README.md")

	if err != nil {
		log.Fatalf("Unable to retrieve chart data from %v - Error: %v", tarballURL, err)
	}

	return readmeContent
}

func getChartTarballURL(repoURL string, chartName string, version string) string {
	
	return utilities.CreateKeyString(repoURL ,"/charts/" ,chartName ,"-" ,version ,".tgz")

}