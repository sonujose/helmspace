package controller

import (
	"log"
	"os"

	"github.com/helm-dimensions/models"
)

// default endpoint for chartmuseum server
var repoURL = "http://localhost:9000"

//GetRepoDetails - details and name of the chart Repo
//TODO: Current Implemtation is limited to only one repo server 
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


