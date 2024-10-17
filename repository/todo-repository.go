package repository

import (
	"database/sql"
	"errors"
	PKG_APP "udit/api-padhai/app"
	"udit/api-padhai/utils"
	// "github.com/gin-gonic/gin"
	// _ "github.com/microsoft/go-mssqldb"
)

var APP_NAME PKG_APP.APP = PKG_APP.TODO_APP

// type repo interface {
// 	GetTodos(ctx *gin.Context) (*sql.Rows, error)
// }

type TodoRepo struct {
	// ctx *gin.Context
}

func (t TodoRepo) CheckIfUserExists(userNameMobileNo string) (bool, string) {
	db := PKG_APP.ConnectToDB(APP_NAME)
	var exists bool
	if db != nil {
		query := "SELECT CASE WHEN EXISTS (SELECT 1 FROM Users WHERE (mobile_no = '" + userNameMobileNo + "' or name = '" + userNameMobileNo + "')) THEN 1 ELSE 0 END"
		db.QueryRow(query).Scan(&exists)
		if exists {
			return true, ""
		} else {
			return false, "User doesn't exist."
		}
	}
	return false, utils.DBError
}

func (t TodoRepo) LoginUser(userNameMobNo string, passWord string) (*sql.Row, error) {
	db := PKG_APP.ConnectToDB(APP_NAME)
	var userId int = -1
	if db != nil {
		query := "SELECT USER_ID FROM Users WHERE (mobile_no = '" + userNameMobNo + "' or name = '" + userNameMobNo + "') and pass = '" + passWord + "'"
		db.QueryRow(query).Scan(&userId)
		if userId != -1 {
			userDataRow, err := t.GetUserDetails(userId)
			if err != nil {
				return nil, err
			} else {
				return userDataRow, nil
			}
		}
	}
	return nil, errors.New(utils.DBError)
}

func (t TodoRepo) GetUserDetails(userid int) (*sql.Row, error) {
	db := PKG_APP.ConnectToDB(APP_NAME)
	if db != nil {
		// result, err := db.ExecContext(ctx, "app_todo_get 1, ''")
		query := "select user_id,name,pass,display_picture,created_on,firebase_token,email_id,mobile_no,is_active,is_premium from users where user_id = ?"
		result := db.QueryRow(query, userid)
		return result, nil
	}
	return nil, errors.New(utils.DBError)
}

func (t TodoRepo) GetTodos(userId string, charStr string) (*sql.Rows, error) {
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
		result, err := db.Query(query,
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
