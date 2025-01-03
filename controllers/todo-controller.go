package controllers

import (
	"database/sql"
	"errors"

	// "encoding/json"

	// "encoding/json"
	"fmt"
	"reflect"
	"strings"

	// "time"
	"udit/api-padhai/models"
	"udit/api-padhai/repository"
	// "udit/api-padhai/utils"
	// "github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
)

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

func getUserDataByUserId(controller TodoController) (*models.User, error) {
	var user models.User = models.User{}
	userDataRow, err := controller.todoRepository.GetUserDetails(1)
	if err != nil {
		return nil, err
	} else {
		userDataRow.Scan(&user.UserID, &user.Name, &user.Password, &user.DisplayPicture, &user.CreatedOn, &user.FirebaseToken, &user.EmailID, &user.MobileNo, &user.IsActive, &user.IsPremium)
		fmt.Println(user)
	}
	return &user, nil
}

func (controller TodoController) TodoApp_userLogin(requestBody models.RequestBodyUserLogin) (*models.User, error) {
	var user models.User = models.User{}
	userNameMobileNo := requestBody.UserNameMobileNo
	passWord := requestBody.Password
	// ipAddress := requestBody.LoginIPAddress
	// platform := requestBody.LoginPlatform
	ifUserExists, errStr := controller.todoRepository.CheckIfUserExists(userNameMobileNo)
	if errStr != "" {
		return nil, errors.New(errStr)
	} else if ifUserExists {
		userDataRow, err := controller.todoRepository.LoginUser(userNameMobileNo, passWord)
		if err != nil {
			return nil, err
		} else {
			userDataRow.Scan(&user.UserID, &user.Name, &user.Password, &user.DisplayPicture, &user.CreatedOn, &user.FirebaseToken, &user.EmailID, &user.MobileNo, &user.IsActive, &user.IsPremium)
			return &user, nil
		}
	}
	return nil, errors.New(errStr)
}

func (controller TodoController) TodoApp_InsertTodoType() {
	result, err := controller.todoRepository.TodoType_Insert()
	if err != nil {
		fmt.Print(err)
		// return nil, errors.New
	} else {
		fmt.Print(result)
	}
}

func (controller TodoController) TodoApp_getTodos(params map[string]string) ([]models.Todo, error) {
	//(*sql.Rows, error) {

	userData, userDataErr := getUserDataByUserId(controller)
	if userDataErr != nil {
		return nil, userDataErr
	}
	fmt.Println(userData)
	var userId = params["userId"]
	var charStr = params["charStr"]
	todoRows, err := controller.todoRepository.GetTodos(userId, charStr)
	if err != nil {
		return nil, err
	} else {
		var todos []models.Todo
		for todoRows.Next() {
			var todo models.Todo
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
