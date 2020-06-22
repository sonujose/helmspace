package models

// Repo - model for helm repo server
type Repo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}