package main

import (
	"fmt"
	"net/http"
	"udit/api-padhai/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func landingPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "startPage.html", nil)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	appServer := gin.Default()
	appServer.LoadHTMLGlob("./html/start_page/*")
	// http://localhost:10000/
	appServer.GET("/", landingPage)

	// "/API/example"
	routes.ExampleRouting(appServer)

	// "/API/todo"
	routes.TodoAppRouting(appServer)

	/* Used for Debugging locally without DOCKER -> go run .*/
	// appServer.Run("localhost:5000")

	/* Used for Debugging locally with DOCKER -> docker run uditnair90/api-padhai-golang:latest */
	// appServer.Run("0.0.0.0:5000")

	/* Used for actual publishing -> -e will tell the port */
	appServer.Run()
}

// docker build . -t uditnair90/api-padhai-golang:latest
// docker run -d -e PORT=10000 -p 10000:10000 uditnair90/api-padhai-golang:latest
// docker image prune -f     ###REMOVES UNUSED IMAGES###
