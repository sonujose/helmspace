<img src="./static/img/icons/apple-touch-icon.png" height="90" width="90">

# Helm-Dimensions

[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)
[![Version](https://badge.fury.io/gh/tterb%2FHyde.svg)](https://badge.fury.io/gh/tterb%2FHyde)

Helm-Dimensions is a web UI dashboard to visualize and manage charts in your Helm repo server. The application is fully written in go language using the gin-gonic web framework.

<div style="display:flex"><img src="./docs/dashboard2.PNG" width="500">
<img src="./docs/details-page.PNG" width="500"></div>

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
