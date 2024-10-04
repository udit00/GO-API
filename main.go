package main

import (
	"net/http"
	"udit/api-padhai/routes"

	"github.com/gin-gonic/gin"
)

// func homePage(w http.ResponseWriter, r *http.Request){
//     fmt.Fprintf(w, "Welcome to the HomePage!")
//     fmt.Println("Endpoint Hit: homePage")
// }

// func handleRequests() {
//     http.HandleFunc("/", homePage)
//     log.Fatal(http.ListenAndServe(":10000", nil))
// }

// func main() {
//     handleRequests()
// }

func main() {
	appServer := gin.Default()
	// appServer.LoadHTMLFiles("./html/start_page/startPage.html")
	appServer.LoadHTMLGlob("./html/start_page/*")

	appServer.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "startPage.html", nil)
	})

	routes.TodoAppRouting(appServer)

	appServer.Run("localhost:10000")
	// connectToDB(app.TODO_APP)
}

// func checkConnection(c *gin.Context) {
// 	query: = sql.
// }

// func connectToDB(appId app.APP) {

// }
