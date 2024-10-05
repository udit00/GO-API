package main

import (
	"net/http"
	"udit/api-padhai/routes"

	"github.com/gin-gonic/gin"
)

func landingPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "startPage.html", nil)
}

func main() {
	appServer := gin.Default()
	appServer.LoadHTMLGlob("./html/start_page/*")
	// http://localhost:10000/
	appServer.GET("/", landingPage)

	// "/API/example"
	routes.ExampleRouting(appServer)

	// "/API/todo"
	routes.TodoAppRouting(appServer)

	appServer.Run("localhost:10000")
}
