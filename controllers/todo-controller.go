package controllers

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"udit/api-padhai/repository"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
)

type Todo struct {
	Todoid      string `json:"todoid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Userid      string `json:"userid"`
	Createdon   string `json:"createdon"`
	// Createid    string `json:"createid"`
	Todotypeid string `json:"todotypeid"`
	Typename   string `json:"typename"`
	Target     string `json:"target"`
}

type User struct {
	Userid         int
	Name           string
	Pass           string
	DisplayPicture string
	IsPremium      string
	CreatedOn      string
	FbToken        string
	EmailId        string
	MobileNo       string
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

func getUserData(controller TodoController) (User, error) {
	var user User = User{}
	userDataRow, error := controller.todoRepository.GetUserDetails(1)
	if error != nil {

	} else {
		fmt.Println(userDataRow)
	}
	return user, nil
}

func (controller TodoController) TodoApp_getTodos(ctx *gin.Context) (*sql.Rows, error) {

	userDataRow, error := getUserData(controller)
	if error != nil {

	} else {
		fmt.Println(userDataRow)
	}
	todoRows, err := controller.todoRepository.GetTodos(ctx)
	if err != nil {
		return nil, err
	} else {

		// var results []Todo
		// for todoRows.Next() {
		// 	var item Todo
		// 	err := todoRows.Scan(&item.Todoid, &item.Title, &item.Description, &item.Userid, &item.Createdon, &item.Todotypeid, &item.Typename, &item.Target)
		// 	if err != nil {
		// 		// ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Row scanning failed"})
		// 		// return
		// 		return nil, err
		// 	} else {
		// 		fmt.Println(&item)
		// 	}
		// 	results = append(results, item)
		// }

		var todos []Todo

		// Scan rows into the slice of Todo
		if err := ScanRows(todoRows, &todos); err != nil {
			return nil, err
		}

		return todoRows, nil
	}
}
