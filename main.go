package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	mirror_router "github.com/chartsmirror/routers"
)

var router *gin.Engine

func main() {
	router := gin.New()

	// Logger middleware will write the logs to gin.DefaultWriter 
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 
	router.Use(gin.Recovery())

	mirror_router.RegisterAPIEndpoints(router)

	router.LoadHTMLGlob("templates/*")

	mirror_router.InitilaizeRoutes(router)

	router.Run(httpPort())
}

func httpPort() string {
	port := "5000"
	if os.Getenv("APP_PORT") != "" {
		port = os.Getenv("APP_PORT")
	}
	return fmt.Sprintf(":%s", port)
}
