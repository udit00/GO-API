package routes

import (
	// "database/sql"

	// "fmt"
	// "io"

	"net/http"

	// PKG_APP "udit/api-padhai/app"
	"udit/api-padhai/controllers"
	"udit/api-padhai/models"
	"udit/api-padhai/utils"

	// "udit/api-padhai/utils"

	// "udit/api-padhai/utils"

	// mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

func TodoAppRouting(router *gin.Engine) {
	// var APP_NAME PKG_APP.APP = PKG_APP.TODO_APP
	todoController := controllers.TodoController{}
	todoController.InitialSetup()
	var todoApiPrefixRoute string = "/API/todo"

	router.GET(todoApiPrefixRoute+"/", func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK, "hello: asd")
		names := [4]string{"udit", "name", "name2", "name3"}
		ctx.JSON(http.StatusOK, names)
	})

	// router.GET(todoApiPrefixRoute+"/getNextTodoTypeID", func(ctx *gin.Context) {
	// 	// ctx.JSON(http.StatusOK, "hello: asd")
	// 	// names := [4]string{"udit", "name", "name2", "name3"}
	// 	// ctx.JSON(http.StatusOK, names)
	// 	todoController.TodoApp_GetTodoTypeNextId(1)
	// 	ctx.JSON(http.StatusAccepted, errors.New("ERROR"))
	// })

	router.POST(todoApiPrefixRoute+"/userLogin", func(ctx *gin.Context) {
		var requestBody models.RequestBodyUserLogin
		if err := ctx.BindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}
		var finalResponse models.ApiResponse
		httpStatus := http.StatusAccepted
		user, err := todoController.TodoApp_userLogin(requestBody)
		if err != nil {
			finalResponse.Status = -1
			finalResponse.Message = err.Error()
		} else {
			finalResponse.Status = 1
			finalResponse.Response = user
		}
		ctx.JSON(httpStatus, finalResponse)

	})

	// router.POST(todoApiPrefixRoute+"/InsertTodoType", func(ctx *gin.Context) {
	// 	var finalResponse models.ApiResponse
	// 	finalResponse.Status = -1
	// 	httpStatus := http.StatusBadRequest
	// 	var postBody models.TodoTypeUpsertPostBodyModel
	// 	if err := ctx.BindJSON(&postBody); err != nil {
	// 		finalResponse.Message = err.Error()
	// 		// panic(finalResponse)
	// 		ctx.JSON(httpStatus, finalResponse)
	// 		return

	// 	}
	// 	if postBody.UserID <= 0 {
	// 		finalResponse.Message = "User Id was not valid."
	// 		ctx.JSON(httpStatus, finalResponse)
	// 		return
	// 	} else if utils.IsNullOrEmpty(postBody.TodoTypeName) {
	// 		finalResponse.Message = "Type Name cannot be Empty."
	// 		ctx.JSON(httpStatus, finalResponse)
	// 		return
	// 	} else if postBody.ColorID <= 0 {
	// 		finalResponse.Message = "Color Id was not valid."
	// 		ctx.JSON(httpStatus, finalResponse)
	// 	}
	// 	httpStatus = http.StatusAccepted
	// 	inserted, responseMsg := todoController.TodoApp_InsertTodoType(postBody)
	// 	if inserted {
	// 		httpStatus = http.StatusOK
	// 		finalResponse.Status = 1
	// 	} else {
	// 		finalResponse.Status = -1
	// 	}
	// 	finalResponse.Message = responseMsg
	// 	ctx.JSON(httpStatus, finalResponse)
	// })

	router.POST(todoApiPrefixRoute+"/UpsertTodo", func(ctx *gin.Context) {
		var finalResponse models.ApiResponse
		finalResponse.Status = -1
		httpStatus := http.StatusBadRequest
		var postBody models.TodoUpsertPostBodyModel
		if err := ctx.BindJSON(&postBody); err != nil {
			finalResponse.Message = err.Error()
			// panic(finalResponse)
			ctx.JSON(httpStatus, finalResponse)
			return

		}
		targetDateTimeProper, timeParsingError := utils.ConvertDateTimeToGoLangTime(postBody.TargetDateTimeString)
		if timeParsingError != nil {
			finalResponse.Status = -1
			finalResponse.Message = timeParsingError.Error()
			ctx.JSON(http.StatusBadRequest, finalResponse)
			return
		}
		postBody.TargetDateTime = targetDateTimeProper
		if postBody.UserID <= 0 {
			finalResponse.Message = "User Id was not valid."
			ctx.JSON(httpStatus, finalResponse)
			return
		} else if utils.IsNullOrEmpty(postBody.Title) {
			finalResponse.Message = "Type Name cannot be Empty."
			ctx.JSON(httpStatus, finalResponse)
			return
		} else if postBody.TodoTypeID <= 0 {
			finalResponse.Message = "Todo Type Id was not valid."
			ctx.JSON(httpStatus, finalResponse)
		}
		httpStatus = http.StatusAccepted
		inserted, responseMsg := todoController.TodoApp_UpsertTodo(postBody)
		if inserted {
			httpStatus = http.StatusOK
			finalResponse.Status = 1
		} else {
			finalResponse.Status = -1
		}
		finalResponse.Message = responseMsg
		ctx.JSON(httpStatus, finalResponse)
	})

	router.POST(todoApiPrefixRoute+"/UpsertTodoType", func(ctx *gin.Context) {
		var finalResponse models.ApiResponse
		finalResponse.Status = -1
		httpStatus := http.StatusBadRequest
		var postBody models.TodoTypeUpsertPostBodyModel
		if err := ctx.BindJSON(&postBody); err != nil {
			finalResponse.Message = err.Error()
			// panic(finalResponse)
			ctx.JSON(httpStatus, finalResponse)
			return

		}
		if postBody.UserID <= 0 {
			finalResponse.Message = "User Id was not valid."
			ctx.JSON(httpStatus, finalResponse)
			return
		} else if utils.IsNullOrEmpty(postBody.TodoTypeName) {
			finalResponse.Message = "Type Name cannot be Empty."
			ctx.JSON(httpStatus, finalResponse)
			return
		} else if postBody.ColorID <= 0 {
			finalResponse.Message = "Color Id was not valid."
			ctx.JSON(httpStatus, finalResponse)
		}
		httpStatus = http.StatusAccepted
		inserted, responseMsg := todoController.TodoApp_UpsertTodoType(postBody)
		if inserted {
			httpStatus = http.StatusOK
			finalResponse.Status = 1
		} else {
			finalResponse.Status = -1
		}
		finalResponse.Message = responseMsg
		ctx.JSON(httpStatus, finalResponse)
	})

	// router.GET(todoApiPrefixRoute+"/userLogin", func(ctx *gin.Context) {
	// 	var userNameMobileNo string = ctx.Query("userNameMobileNo")
	// 	var passWord string = ctx.Query("passWord")
	// 	params := make(map[string]string)
	// 	params["userNameMobileNo"] = userNameMobileNo
	// 	params["passWord"] = passWord
	// 	var finalResponse models.ApiResponse
	// 	httpStatus := http.StatusAccepted
	// 	user, err := todoController.TodoApp_userLogin(params)
	// 	if err != nil {
	// 		finalResponse.Status = -1
	// 		finalResponse.Message = err.Error()
	// 	} else {
	// 		finalResponse.Status = 1
	// 		finalResponse.Response = user
	// 	}
	// 	ctx.JSON(httpStatus, finalResponse)
	// })

	router.GET(todoApiPrefixRoute+"/getTodos", func(ctx *gin.Context) {
		// fmt.Println(ctx.Query("userId"))
		// fmt.Println(ctx.Request.URL.Query())
		var userId string = ctx.Query("userId")
		var charStr string = ctx.Query("charStr")

		params := make(map[string]string)
		params["userId"] = userId
		params["charStr"] = charStr
		var finalResponse models.ApiResponse
		httpStatus := http.StatusAccepted
		todos, err := todoController.TodoApp_getTodos(params)
		if err != nil {
			finalResponse.Status = -1
			finalResponse.Message = err.Error()
		} else {
			finalResponse.Status = 1
			finalResponse.Response = todos
		}
		ctx.JSON(httpStatus, finalResponse)
	})

	// router.PATCH(todoApiPrefixRoute+"/updateFireBaseToken", func(ctx *gin.Context) {
	// 	var userId string = ctx.Query("userId")
	// 	var fbToken string = ctx.Query("firebaseToken")

	// 	params := make
	// })

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
