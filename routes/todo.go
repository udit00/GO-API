package routes

import (
	"database/sql"
	"net/http"
	PKG_APP "udit/api-padhai/app"
	"udit/api-padhai/models"
	"udit/api-padhai/utils"

	// mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

var APP_NAME PKG_APP.APP = PKG_APP.TODO_APP

func connectToDB(application PKG_APP.APP) *sql.DB {
	connString := PKG_APP.GetDbConnString(application)
	db, err := sql.Open("mssql", connString)
	if err != nil {
		return nil
	}
	return db
}

func TodoAppRouting(router *gin.Engine) {
	// var APP_NAME PKG_APP.APP = PKG_APP.TODO_APP
	var todoApiPrefixRoute string = "/API/todo"

	router.GET(todoApiPrefixRoute+"/", func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK, "hello: asd")
		names := [4]string{"udit", "name", "name2", "name3"}
		ctx.JSON(http.StatusOK, names)
	})

	router.GET(todoApiPrefixRoute+"/getTodos", getTodosWithSP)

}

func getTodos(ctx *gin.Context) {
	var finalResponse models.ApiResponse
	httpStatus := http.StatusAccepted
	db := connectToDB(APP_NAME)
	// fmt.Println(db.Stats().OpenConnections)
	if db != nil {
		rows, err := db.Query("select * from todo;")
		if err != nil {
			finalResponse = utils.GetErrorResponse(err.Error(), err)
		} else {
			json, err := utils.ReturnJsonFromRows(rows)
			if err != nil {
				finalResponse = utils.GetErrorResponse(err.Message, err)
			} else {
				httpStatus = http.StatusOK
				finalResponse = utils.GetSuccessResponse(json)
			}
		}
	} else {
		finalResponse = utils.GetErrorResponse("Something went wrong.", nil)
	}
	ctx.JSON(httpStatus, finalResponse)
}

func getTodosWithSP(ctx *gin.Context) {
	var finalResponse models.ApiResponse
	httpStatus := http.StatusAccepted
	db := connectToDB(APP_NAME)
	if db != nil {

		// result, err := db.ExecContext(ctx, "app_todo_get 1, ''")
		execQuery := "app_todo_get @prm_userid=?, @prm_searchstr=?"
		result, err := db.QueryContext(ctx, execQuery,
			sql.NamedArg{
				Name:  "p1",
				Value: 1,
			},
			sql.NamedArg{
				Name:  "p2",
				Value: "",
			},
		)

		// sql.Named("prm_userid", sql),
		// sql.Named("prm_charstr", ""),

		if err != nil {
			finalResponse = utils.GetErrorResponse(err.Error(), err)
		} else {
			// ctx.JSON(http.StatusOK, result)
			rows, err := utils.ReturnJsonFromRows(result)
			if err != nil {
				finalResponse = utils.GetErrorResponse(err.Error(), err)
			} else {
				finalResponse = utils.GetSuccessResponse(rows)
			}
		}
	} else {
		finalResponse = utils.GetErrorResponse("Something went wrong.", nil)
	}
	ctx.JSON(httpStatus, finalResponse)
}
