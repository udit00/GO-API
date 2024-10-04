package routes

import (
	"database/sql"
	"net/http"
	PKG_APP "udit/api-padhai/app"
	"udit/api-padhai/models"

	mssql "github.com/denisenkom/go-mssqldb"
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

func returnJsonFromRows(rows *sql.Rows) (fJson any, fError *mssql.Error) {
	columnTypes, err := rows.ColumnTypes()

	if err != nil {
		return nil, &mssql.Error{Message: err.Error()}
	}

	count := len(columnTypes)
	finalRows := []interface{}{}

	for rows.Next() {

		scanArgs := make([]interface{}, count)

		for i, v := range columnTypes {

			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
				break
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
				break
			case "INT4":
				scanArgs[i] = new(sql.NullInt64)
				break
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := rows.Scan(scanArgs...)

		if err != nil {
			return nil, &mssql.Error{Message: err.Error()}
		}

		masterData := map[string]interface{}{}

		for i, v := range columnTypes {

			if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
				masterData[v.Name()] = z.Bool
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullString); ok {
				masterData[v.Name()] = z.String
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
				masterData[v.Name()] = z.Int64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
				masterData[v.Name()] = z.Float64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
				masterData[v.Name()] = z.Int32
				continue
			}

			masterData[v.Name()] = scanArgs[i]
		}

		finalRows = append(finalRows, masterData)
	}
	return finalRows, nil
	// z, err := json.Marshal(finalRows)
	// if err != nil {
	// 	return nil, &mssql.Error{Message: err.Error()}
	// } else {
	// 	return string(z), nil
	// }
	// return z, nil
}

func TodoAppRouting(router *gin.Engine) {
	// var APP_NAME PKG_APP.APP = PKG_APP.TODO_APP
	var todoApiPrefixRoute string = "/API/todo"

	router.GET(todoApiPrefixRoute+"/", func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK, "hello: asd")
		names := [4]string{"udit", "name", "name2", "name3"}
		ctx.JSON(http.StatusOK, names)
	})

	router.GET(todoApiPrefixRoute+"/getTodos", getTodos)

}

func getTodos(ctx *gin.Context) {
	finalResponse := models.ApiResponse{Status: -1, Message: ""}
	httpStatus := http.StatusAccepted
	db := connectToDB(APP_NAME)
	// fmt.Println(db.Stats().OpenConnections)
	if db != nil {
		rows, err := db.Query("select * from todo;")
		if err != nil {
			finalResponse.Message = err.Error()
			finalResponse.Response = err
		} else {
			json, err := returnJsonFromRows(rows)
			if err != nil {
				finalResponse.Message = err.Message
				finalResponse.Response = err
			} else {
				httpStatus = http.StatusOK
				finalResponse.Status = 1
				finalResponse.Response = json
			}
		}
	} else {
		finalResponse.Response = "Something went wrong"
	}
	ctx.JSON(httpStatus, finalResponse)
}
