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
		query := "select user_id,name,pass,display_picture,created_on,firebase_token,email_id,mobile_no,is_active,is_premium from users where user_id = ?"
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

func (t TodoRepo) GetTodos(ctx *gin.Context, userId string, charStr string) (*sql.Rows, error) {
	db := PKG_APP.ConnectToDB(APP_NAME)
	if db != nil {
		query := `select  	t.todo_id, t.title, t.description,
							u.name, t.created_on, t.target,
							tt.type_id, tt.type_name
						from todo t
						inner join todo_type tt on tt.type_id = t.type_id
						inner join users u on u.user_id = t.user_id
						where t.user_id = ?
						and isnull(t.title, '') like '%' + ? + '%'`

		// query := "exec app_todo_get @prm_userid = ?"
		result, err := db.QueryContext(ctx, query,
			sql.NamedArg{
				Name:  "p1",
				Value: userId,
			},

			sql.NamedArg{
				Name:  "p2",
				Value: charStr,
			},
		)
		if err != nil {
			return nil, err
		} else {
			return result, nil
		}
	}
	return nil, errors.New(utils.DBError)
}
