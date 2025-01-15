package PKG_APP

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strconv"
	// "udit/api-padhai/utils"
)

type APP string

const (
	TODO_APP  APP = "todo"
	EZONE_APP APP = "ezone"
)

func ConnectToDB(application APP) *sql.DB {
	connString := GetDbConnString(application)
	db, err := sql.Open("mssql", connString)
	if err != nil {
		return nil
	}
	return db
}

func GetDbConnString(app APP) string {
	convertedPort, errorWhileConvertPortToInt := strconv.Atoi(os.Getenv("DB_PORT"))
	if errorWhileConvertPortToInt != nil {
		return ""
	}
	var host = os.Getenv("DB_HOST")
	var port = convertedPort
	var userId string = os.Getenv("DB_USER")
	var passWord string = os.Getenv("DB_PASSWORD")
	var connString string
	var app_name string
	switch app {
	case TODO_APP:
		app_name = "todo"
	case EZONE_APP:
		app_name = "ezone"
	}
	query := url.Values{}
	// query.Add("app name", app_name)
	query.Add("database", app_name)
	query.Add("encrypt", "disable")
	// connString = "sqlserver://sa:GreenTrans@123*@" + host + "?database=" + app_name + "&connection+timeout=30"
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(userId, passWord),
		Host:     fmt.Sprintf("%s:%d", host, port),
		RawQuery: query.Encode(),
	}
	connString = u.String()
	return connString
}
