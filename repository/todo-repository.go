package repository

import (
	"database/sql"
	"errors"
	"fmt"

	// "internal/itoa"
	// "strconv"

	PKG_APP "udit/api-padhai/app"
	"udit/api-padhai/models"
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

func (t TodoRepo) CheckIfRowExists(query string) (doesTheRowExists bool, err error) {
	db := PKG_APP.ConnectToDB(APP_NAME)
	var exists bool
	if db != nil {
		query := "SELECT CASE WHEN EXISTS (" + query + ") THEN CAST(1 AS BIT) ELSE CAST(0 AS BIT) END"
		// db.QueryRow(query).Scan(&exists)
		result := db.QueryRow(query)
		err := result.Err()
		if err != nil {
			fmt.Print(err)
			return false, err
		} else {
			result.Scan(&exists)
			return exists, nil
		}
	}
	return false, errors.New("server error")
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

func (t TodoRepo) GetTodos(userId int, charStr string) (*sql.Rows, error) {
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

// func (t TodoRepo) TodoType_Update() {
// 	db := PKG_APP.ConnectToDB(APP_NAME)
// 	if db != nil {
// 		query := `select t.todoid, t.title,
// 					u.name, u.firebase_token
// 					from todo t
// 					inner join users u on u.user_id = t.user_id
// 					where t.target
// 					`
// 	}
// }

func (t TodoRepo) NextTodoIDAsPerUser(userId int) (*sql.Row, error) {
	db := PKG_APP.ConnectToDB(APP_NAME)
	if db != nil {
		query := `Select isnull(max(todo_id), 0) + 1 from todo where user_id = ?`
		result := db.QueryRow(query,
			sql.NamedArg{
				Name:  "p1",
				Value: userId,
			},
		)
		if result != nil {
			return result, nil
		}
		return nil, result.Err()
	}
	return nil, errors.New(utils.DBError)
}

func (t TodoRepo) Todo_Insert(todoModel models.TodoUpsertPostBodyModel) (*sql.Rows, error) {
	currentTime := utils.GetCurrentDateTimeForSqlString()
	db := PKG_APP.ConnectToDB(APP_NAME)
	if db != nil {
		// currentTime := time.Now()
		query := `Insert into todo(todo_id, title, description, user_id, created_on, type_id, target, completion_status_id)
					values(?, ?, ?, ?, ?, ?, ?, ?)`
		result, err := db.Query(query,
			sql.NamedArg{
				Name:  "p1",
				Value: todoModel.TodoID,
			},
			sql.NamedArg{
				Name:  "p2",
				Value: todoModel.Title,
			},
			sql.NamedArg{
				Name:  "p3",
				Value: todoModel.Description,
			},
			sql.NamedArg{
				Name:  "p4",
				Value: todoModel.UserID,
				// Value: currentTime.Format(utils.DATE_TIME_DEFAULT_FORMAT),
			},
			sql.NamedArg{
				Name:  "p5",
				Value: currentTime,
			},
			sql.NamedArg{
				Name:  "p6",
				Value: todoModel.TodoTypeID,
			},
			sql.NamedArg{
				Name:  "p7",
				Value: todoModel.TargetDateTime,
			},
			sql.NamedArg{
				Name:  "p8",
				Value: todoModel.CompletionStatusID,
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

func (t TodoRepo) NextTodoTypeIDAsPerUser(userId int) (*sql.Row, error) {
	db := PKG_APP.ConnectToDB(APP_NAME)
	if db != nil {
		query := `Select isnull(max(TYPE_ID), 0) + 1 from todo_type where create_id = ?`
		result := db.QueryRow(query,
			sql.NamedArg{
				Name:  "p1",
				Value: userId,
			},
		)
		if result != nil {
			return result, nil
		}
		return nil, result.Err()
	}
	return nil, errors.New(utils.DBError)
}

func (t TodoRepo) TodoType_Insert(todoType models.TodoTypeModel) (*sql.Rows, error) {
	currentTime := utils.GetCurrentDateTimeForSqlString()
	db := PKG_APP.ConnectToDB(APP_NAME)
	if db != nil {
		// currentTime := time.Now()
		query := `Insert into todo_type(type_id, type_name, create_id, created_on, color_id)
					values(?, ?, ?, ?, ?)`
		result, err := db.Query(query,
			sql.NamedArg{
				Name:  "p1",
				Value: todoType.TypeId,
			},
			sql.NamedArg{
				Name:  "p2",
				Value: todoType.TypeName,
			},
			sql.NamedArg{
				Name:  "p3",
				Value: todoType.CreateID,
			},
			sql.NamedArg{
				Name:  "p4",
				Value: currentTime,
				// Value: currentTime.Format(utils.DATE_TIME_DEFAULT_FORMAT),
			},
			sql.NamedArg{
				Name:  "p5",
				Value: todoType.ColorID,
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

func (t TodoRepo) TodoType_Update(todoType models.TodoTypeModel) (*sql.Rows, error) {
	// currentTime := utils.GetCurrentDateTimeForSqlString()
	db := PKG_APP.ConnectToDB(APP_NAME)
	if db != nil {
		// currentTime := time.Now()

		query := `Update todo_type set type_name = ?, color_id = ? where type_id = ` + utils.ConvertIntToString(todoType.TypeId) + `and create_id = ` + utils.ConvertIntToString(todoType.CreateID)
		// query := `Insert into todo_type(type_id, type_name, create_id, created_on, color_id)
		// 			values(?, ?, ?, ?, ?)`
		result, err := db.Query(query,
			sql.NamedArg{
				Name:  "p1",
				Value: todoType.TypeName,
			},
			sql.NamedArg{
				Name:  "p2",
				Value: todoType.ColorID,
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
