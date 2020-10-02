# <img src="./static/img/icons/apple-touch-icon.png" height="25" width="25"> Helmspace
[![Build Status](https://dev.azure.com/sonujse/Helmspace/_apis/build/status/helmspace?branchName=master)](https://dev.azure.com/sonujse/Helmspace/_build/latest?definitionId=1&branchName=master)
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)
[![Version](https://badge.fury.io/gh/tterb%2FHyde.svg)](https://badge.fury.io/gh/tterb%2FHyde)

Helmspace is a web UI dashboard to visualize and manage charts in your Helm repository. 

## Application Usage

The application is available as docker image so that you can host it in your premise and connect to any public or private Chartmuseum server.

The Docker image available at [bhsonu/helmspace](https://hub.docker.com/r/bhsonu/helmspace)

Chartmusum server url is required to run the application. Refer to [Chartmuseum](https://github.com/helm/chartmuseum) to setup Chartmuseum server

```
docker run --rm -it \
  -p 5000:5000 \
  -e CHART_MUSEUM_URL=http://localhost:9000
  bhsonu/helmspace:latest
```
After successfull docker run, the dashboard will be available at `http://localhost:5000`  

## Development

The application is fully written in go language using the gin-gonic web framework.

* [gin-gonic](https://gin-gonic.com/) - The web framework for GO
* [go](https://golang.org/) - Programing language
* [go-modules](https://github.com/golang/go/wiki/Modules) - Go Package management

You need go 1.11+ installed in the machine for building the application without Docker. 

Use `docker-compose.yaml` to build and run locally

## Configuration

| Parameter          | Default                  | Description          |
| ------------------ | ------------------------ | ---------------------|
|CHART_MUSEUM_URL    | http://localhost:9000    |  URL Endpoint of your chartmuseum server |
|CHART_MUSEUM_API    | /api/charts              |  Chartmuseum API Endpoint                |
|APP_PORT            | 5000                     |  Application Port                        |

## Limitations
The current version only supports Chartmuseum Helm server, future release will support all sorts of Helm servers
