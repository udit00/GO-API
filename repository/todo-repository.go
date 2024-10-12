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
		query := "select userid, name, pass, isnull(displaypicture, 'x'), ispremium, createdon, fbtoken, email_id, mobile_no from users where userid = ?"
		result := db.QueryRow(query, userid) // sql.NamedArg{
		// 	Name:  "p1",
		// 	Value: 1,
		// },
		// sql.NamedArg{
		// 	Name:  "p2",
		// 	Value: "",
		// },

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
