package controllers

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	// "encoding/json"

	// "encoding/json"
	"fmt"
	"reflect"
	"strings"

	// "time"
	"udit/api-padhai/models"
	"udit/api-padhai/repository"
	"udit/api-padhai/utils"
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

func (controller TodoController) getUserDataByUserId(userId int) (*models.User, error) {
	var user models.User = models.User{}
	userDataRow, err := controller.todoRepository.GetUserDetails(userId)
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
	checkIfUserExistsQuery := "SELECT 1 FROM Users WHERE (mobile_no = '" + userNameMobileNo + "' or name = '" + userNameMobileNo + "')"
	ifUserExists, errFromIfUserExists := controller.todoRepository.CheckIfRowExists(checkIfUserExistsQuery)
	if errFromIfUserExists != nil {
		return nil, errFromIfUserExists
	} else if ifUserExists {
		return nil, errors.New("user doesn't exist")
	} else {
		userDataRow, err := controller.todoRepository.LoginUser(userNameMobileNo, passWord)
		if err != nil {
			return nil, err
		} else {
			userDataRow.Scan(&user.UserID, &user.Name, &user.Password, &user.DisplayPicture, &user.CreatedOn, &user.FirebaseToken, &user.EmailID, &user.MobileNo, &user.IsActive, &user.IsPremium)
			return &user, nil
		}
	}
	// return nil, errors.New(errStr)
}

func (controller TodoController) TodoApp_GetNextTodoId(userId int) int {
	queryRow, err := controller.todoRepository.NextTodoIDAsPerUser(userId)
	if err != nil {
		fmt.Print(err)
		return -1
	}
	nextTodoTypeIdForUser := 0
	queryRow.Scan(&nextTodoTypeIdForUser)
	return nextTodoTypeIdForUser
}

func (controller TodoController) TodoApp_UpsertTodo(postBody models.TodoUpsertPostBodyModel) (bool, string) {
	var todoExists = false
	if postBody.TodoID > 0 {
		checkIfTodoExistsQuery := "Select 1 from todo where todo_id = " + utils.ConvertIntToString(postBody.TodoID) + " and user_id = " + utils.ConvertIntToString(postBody.UserID)
		var errFromIfTodoExists error
		todoExists, errFromIfTodoExists = controller.todoRepository.CheckIfRowExists(checkIfTodoExistsQuery)
		if errFromIfTodoExists != nil {
			return false, errFromIfTodoExists.Error()
		}
	}
	if todoExists {
		return false, "Yet to implement."
	} else {
		nextTodoID := controller.TodoApp_GetNextTodoId(postBody.UserID)
		if nextTodoID <= 0 {
			return false, "Could not create todo."
		} else {
			postBody.TodoID = nextTodoID
			_, insertTodoErr := controller.todoRepository.Todo_Insert(postBody)
			if insertTodoErr != nil {
				return false, insertTodoErr.Error()
			}
			return true, "Todo was saved successfully."
		}
	}
	// return false, "Something went wrong"
}

func (controller TodoController) TodoApp_GetNextTodoTypeId(userId int) int {
	queryRow, err := controller.todoRepository.NextTodoTypeIDAsPerUser(userId)
	if err != nil {
		fmt.Print(err)
	}
	nextTodoTypeIdForUser := 0
	queryRow.Scan(&nextTodoTypeIdForUser)
	return nextTodoTypeIdForUser
}

func (controller TodoController) TodoApp_UpsertTodoType(postBody models.TodoTypeUpsertPostBodyModel) (bool, string) {
	checkIfTodoTypeExistsQuery := "select 1 from todo_type where type_name = ltrim(rtrim('" + postBody.TodoTypeName + "')) and create_id = " + utils.ConvertIntToString(postBody.UserID)
	todoTypeExistsWithName, err := controller.todoRepository.CheckIfRowExists(checkIfTodoTypeExistsQuery)
	if err != nil {
		return false, err.Error()
	} else if todoTypeExistsWithName {
		return false, "Todo type already exists."
	}
	if postBody.TodoTypeID <= 0 {
		// checkIfTodoTypeExistsQuery := "select 1 from todo_type where type_name = ltrim(rtrim('" + postBody.TodoTypeName + "'))"
		// rowExists := controller.todoRepository.CheckIfRowExists(checkIfTodoTypeExistsQuery)
		todoTypeIdToInsert := controller.TodoApp_GetNextTodoTypeId(postBody.UserID)
		todoTypeModel := models.TodoTypeModel{
			TypeId:    todoTypeIdToInsert,
			TypeName:  postBody.TodoTypeName,
			CreateID:  postBody.UserID,
			CreatedOn: time.Now(),
			ColorID:   postBody.ColorID,
		}
		_, err := controller.todoRepository.TodoType_Insert(todoTypeModel)

		if err != nil {
			return false, err.Error()
			// return nil, errors.New
		} else {
			successMsg := "Todo Type successfully saved."
			return true, successMsg
		}
	} else {
		checkIfTodoTypeExistsQuery := "select 1 from todo_type where type_id = " + utils.ConvertIntToString(postBody.TodoTypeID) + " and create_id = " + utils.ConvertIntToString(postBody.UserID)
		todoTypeExistsWithIDForUpdate, errFromIfTheTodoTypeExists := controller.todoRepository.CheckIfRowExists(checkIfTodoTypeExistsQuery)
		if errFromIfTheTodoTypeExists != nil {
			return false, errFromIfTheTodoTypeExists.Error()
		} else if todoTypeExistsWithIDForUpdate {
			todoTypeModel := models.TodoTypeModel{
				TypeId:    postBody.TodoTypeID,
				TypeName:  postBody.TodoTypeName,
				CreateID:  postBody.UserID,
				CreatedOn: time.Now(),
				ColorID:   postBody.ColorID,
			}
			_, err := controller.todoRepository.TodoType_Update(todoTypeModel)
			if err != nil {
				return false, err.Error()
				// return nil, errors.New
			} else {
				successMsg := "Todo Type Updated Successfully."
				return true, successMsg
			}
		} else {
			return false, "Todo type does not exists for Type Id# " + utils.ConvertIntToString(postBody.TodoTypeID)
		}

	}
}

func (controller TodoController) TodoApp_getTodos(params map[string]string) ([]models.Todo, error) {
	//(*sql.Rows, error) {

	var strUserId = params["userId"]
	var charStr = params["charStr"]
	userId, errParsingStrToInt := strconv.Atoi(strUserId)
	if errParsingStrToInt != nil {
		return nil, errParsingStrToInt
	}
	userData, userDataErr := controller.getUserDataByUserId(userId)
	if userDataErr != nil {
		return nil, userDataErr
	}
	fmt.Println(userData)
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
