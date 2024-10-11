package repository

import (
	"database/sql"
	"errors"
	PKG_APP "udit/api-padhai/app"
	"udit/api-padhai/utils"

	"github.com/gin-gonic/gin"
	// _ "github.com/microsoft/go-mssqldb"
)

var APP_NAME PKG_APP.APP = PKG_APP.TODO_APP

// type repo interface {
// 	GetTodos(ctx *gin.Context) (*sql.Rows, error)
// }

type TodoRepo struct {
	// ctx *gin.Context
}

func (t TodoRepo) GetUserDetails(userid int) (*sql.Row, error) {
	db := PKG_APP.ConnectToDB(APP_NAME)
	if db != nil {
		// result, err := db.ExecContext(ctx, "app_todo_get 1, ''")
		query := "SELECT USERID, NAME, PASS, DISPLAYPICTURE, ISPREMIUM, CREATEDON, FBTOKEN, EMAIL_ID, MOBILE_NO FROM USERS"
		result := db.QueryRow(query,
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

		// if err != nil {
		// 	return nil, err
		// } else {
		// 	return result, nil
		// }
		return result, nil
	}
	return nil, errors.New(utils.DBError)
}

func (t TodoRepo) GetTodos(ctx *gin.Context) (*sql.Rows, error) {
	db := PKG_APP.ConnectToDB(APP_NAME)
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
			return nil, err
		} else {
			return result, nil
		}
	}
	return nil, errors.New(utils.DBError)
}
