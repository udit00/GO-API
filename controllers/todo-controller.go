package controllers

import (
	"database/sql"
	// "encoding/json"

	// "encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
	"udit/api-padhai/repository"

	// "udit/api-padhai/utils"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
)

type Todo struct {
	TodoID      int    `json:"todo_id" db:"todo_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	UserName    string `json:"name" db:"name"` // Assuming this is the user's name
	CreatedOn   string `json:"created_on" db:"created_on"`
	Target      string `json:"target" db:"target"`
	TypeID      int    `json:"type_id" db:"type_id"`
	TypeName    string `json:"type_name" db:"type_name"`
}

type User struct {
	UserID         int       `db:"user_id"`
	Name           string    `db:"name"`
	Password       string    `db:"pass"` // Avoid using "pass" as it can be misleading. Consider using "Password" instead.
	DisplayPicture *string   `db:"display_picture"`
	CreatedOn      time.Time `db:"created_on"`
	FirebaseToken  *string   `db:"firebase_token"`
	EmailID        *string   `db:"email_id"`
	MobileNo       string    `db:"mobile_no"`
	IsActive       bool      `db:"is_active"`
	IsPremium      bool      `db:"is_premium"`
}

type TodoController struct {
	// ctx *gin.Context
	todoRepository repository.TodoRepo
}

func NewController() *TodoController {
	return &TodoController{todoRepository: repository.TodoRepo{}}
}

func ScanRows(rows *sql.Rows, result interface{}) error {
	// Get the slice value
	v := reflect.ValueOf(result)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("result must be a pointer to a slice")
	}

	// Get the type of T
	elemType := v.Elem().Type().Elem()

	// Get the columns from rows
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// Create a slice to hold the values
	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}

	// Iterate through the rows
	for rows.Next() {
		// Scan the row into values
		if err := rows.Scan(values...); err != nil {
			return err
		}

		// Create a new struct of type T
		item := reflect.New(elemType).Elem()

		// Populate the struct fields
		for i, colName := range columns {
			// Use the field tag to find the corresponding struct field
			field := item.FieldByNameFunc(func(name string) bool {
				// Compare the field name with the tag
				field, _ := elemType.FieldByName(name)
				return strings.EqualFold(field.Tag.Get("db"), colName)
			})

			if field.IsValid() && field.CanSet() {
				val := *(values[i].(*interface{}))
				field.Set(reflect.ValueOf(val))
			}
		}

		// Append the item to the result slice
		v.Elem().Set(reflect.Append(v.Elem(), item))
	}

	return nil
}

func getUserData(controller TodoController) (*User, error) {
	var user User = User{}
	userDataRow, err := controller.todoRepository.GetUserDetails(1)
	if err != nil {
		return nil, err
	} else {
		userDataRow.Scan(&user.UserID, &user.Name, &user.Password, &user.DisplayPicture, &user.CreatedOn, &user.FirebaseToken, &user.EmailID, &user.MobileNo, &user.IsActive, &user.IsPremium)
		fmt.Println(user)
	}
	return &user, nil
}

func (controller TodoController) TodoApp_getTodos(ctx *gin.Context, params map[string]string) ([]Todo, error) {
	//(*sql.Rows, error) {

	userData, userDataErr := getUserData(controller)
	if userDataErr != nil {
		return nil, userDataErr
	}
	fmt.Println(userData)
	var userId = params["userId"]
	var charStr = params["charStr"]
	todoRows, err := controller.todoRepository.GetTodos(ctx, userId, charStr)
	if err != nil {
		return nil, err
	} else {
		var todos []Todo
		for todoRows.Next() {
			var todo Todo
			if err := todoRows.Scan(&todo.TodoID, &todo.Title, &todo.Description, &todo.UserName, &todo.CreatedOn, &todo.Target, &todo.TypeID, &todo.TypeName); err != nil {
				// handle error
			}
			todos = append(todos, todo)
		}

		if err := todoRows.Err(); err != nil {
			// handle error
		}

		// for _, todo := range todos {
		// 	fmt.Printf("Todo ID: %d\n", todo.TodoID)
		// 	fmt.Printf("Title: %s\n", todo.Title)
		// 	fmt.Printf("Description: %s\n", todo.Description)
		// 	fmt.Printf("User: %s\n", todo.UserName)
		// 	fmt.Printf("Created On: %s\n", todo.CreatedOn)
		// 	fmt.Printf("Target: %s\n", todo.Target)
		// 	fmt.Printf("Type ID: %d\n", todo.TypeID)
		// 	fmt.Printf("Type Name: %s\n", todo.TypeName)
		// 	fmt.Println()
		// }
		return todos, nil
	}
}
