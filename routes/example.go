package routes

import (
	// "database/sql"
	// "database/sql"
	"net/http"
	"udit/api-padhai/utils"

	// "udit/api-padhai/utils"

	// PKG_APP "udit/api-padhai/app"

	// mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

func ExampleRouting(router *gin.Engine) {
	// var APP_NAME PKG_APP.APP = PKG_APP.TODO_APP
	var exampleApiPrefixRoute string = "/API/example"

	router.GET(exampleApiPrefixRoute+"/", func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK, "hello: asd")
		utils.GetCurrentDateTimeForSqlString()
		names := [4]string{"udit", "name", "name2", "name3"}
		ctx.JSON(http.StatusOK, names)
	})

	// router.GET(exampleApiPrefixRoute+"/callSPWIthParams_ExecContext", func(ctx *gin.Context) {
	// 	var fResult any
	// 	db := connectToDB(APP_NAME)
	// 	result, err := db.ExecContext(ctx, "app_todo_get 1, ''")
	// 	if err != nil {
	// 		fResult = err
	// 	} else {
	// 		fResult = result
	// 	}
	// 	ctx.JSON(http.StatusOK, fResult)
	// })

	// router.GET(exampleApiPrefixRoute+"/callSPWIthParams_QueryContext", func(ctx *gin.Context) {
	// 	var fResult any
	// 	db := connectToDB(APP_NAME)
	// 	execQuery := "app_todo_get @prm_userid=?, @prm_searchstr=?"
	// 	result, err := db.QueryContext(ctx, execQuery,
	// 		sql.NamedArg{
	// 			Name:  "p1",
	// 			Value: 1,
	// 		},
	// 		sql.NamedArg{
	// 			Name:  "p2",
	// 			Value: "",
	// 		},
	// 	)
	// 	if err != nil {
	// 		fResult = err
	// 	} else {
	// 		rows, _ := utils.ReturnJsonFromRows(result)
	// 		fResult = rows
	// 	}
	// 	ctx.JSON(http.StatusOK, fResult)
	// })

}
