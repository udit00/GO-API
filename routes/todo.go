package routes

import (
	// "database/sql"
	"net/http"
	// PKG_APP "udit/api-padhai/app"
	"udit/api-padhai/controllers"
	"udit/api-padhai/models"

	// "udit/api-padhai/utils"

	// "udit/api-padhai/utils"

	// mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

func TodoAppRouting(router *gin.Engine) {
	// var APP_NAME PKG_APP.APP = PKG_APP.TODO_APP
	todoController := controllers.TodoController{}
	var todoApiPrefixRoute string = "/API/todo"

	router.GET(todoApiPrefixRoute+"/", func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK, "hello: asd")
		names := [4]string{"udit", "name", "name2", "name3"}
		ctx.JSON(http.StatusOK, names)
	})

	router.GET(todoApiPrefixRoute+"/getTodos", func(ctx *gin.Context) {
		var finalResponse models.ApiResponse
		httpStatus := http.StatusAccepted
		todos, err := todoController.TodoApp_getTodos(ctx)
		if err != nil {
			finalResponse.Status = -1
			finalResponse.Message = err.Error()
		} else {
			finalResponse.Status = 1
			// json, err := utils.ReturnJsonFromRows(todos)
			if err != nil {
				finalResponse.Status = -1
				finalResponse.Response = err
			} else {
				finalResponse.Response = todos
			}
		}
		ctx.JSON(httpStatus, finalResponse)
	})

}

// func getTodos(ctx *gin.Context) {
// 	var finalResponse models.ApiResponse
// 	httpStatus := http.StatusAccepted
// 	db := connectToDB(APP_NAME)
// 	// fmt.Println(db.Stats().OpenConnections)
// 	if db != nil {
// 		rows, err := db.Query("select * from todo;")
// 		if err != nil {
// 			finalResponse = utils.GetErrorResponse(err.Error(), err)
// 		} else {
// 			json, err := utils.ReturnJsonFromRows(rows)
// 			if err != nil {
// 				finalResponse = utils.GetErrorResponse(err.Message, err)
// 			} else {
// 				httpStatus = http.StatusOK
// 				finalResponse = utils.GetSuccessResponse(json)
// 			}
// 		}
// 	} else {
// 		finalResponse = utils.GetErrorResponse("Something went wrong.", nil)
// 	}
// 	ctx.JSON(httpStatus, finalResponse)
// }
