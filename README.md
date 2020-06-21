# Helm-Dimensions

Helm-Dimensions is a web UI dashboard to visualize and manage charts in your Helm repo server. The application is fully written in go language using the gin-gonic web framework.

<img src="./docs/dashboard2.PNG">

## Limitations
The current version only supports Chartmuseum Helm server, future release will support all sorts of Helm servers

## Build with

* [gin-gonic](https://gin-gonic.com/) - The web framework for GO
* [go](https://golang.org/) - Programing language

## Working Locally with Chartmuseum Server
```
# Provide the chartmuseum server url
setx CHART_MUSEUM_URL "http://my-chartmuseum-server"

go build

./helm-dimensions
```
