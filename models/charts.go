package models

import (
	"encoding/json"
	"time"
)

// Chart reprecents a helm chart with all avaible info
type Chart struct {
	Name        string       `json:"name"`
	Home        string       `json:"home"`
	Sources     []string     `json:"sources"`
	Version     string       `json:"version"`
	AppVersion  string       `json:"appVersion"`
	Description string       `json:"description"`
	Maintainers []Maintainer `json:"maintainers"`
	Engine      string       `json:"engine"`
	Icon        string       `json:"icon"`
	Urls        []string     `json:"urls"`
	Created     time.Time    `json:"created"`
	Digest      string       `json:"digest"`
}

// Maintainer reprecents a chart maintainer
type Maintainer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ChartItem - chart name
type ChartItem struct {
	Name string `uri:"name" binding:"required"`
}

// GetNewCharts - unmarshall data
func GetNewCharts(data []byte) (map[string][]Chart, error) {
	var c map[string][]Chart
	err := json.Unmarshal(data, &c)
	return c, err
}

// GetNewChartItem - unmarshall data
func GetNewChartItem(data []byte) ([]Chart, error) {
	var c []Chart
	err := json.Unmarshal(data, &c)
	return c, err
}