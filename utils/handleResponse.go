package utils

import (
	"database/sql"
	"udit/api-padhai/models"

	mssql "github.com/denisenkom/go-mssqldb"
)

func GetSuccessResponse(response any) models.ApiResponse {
	return models.ApiResponse{Status: 1, Message: "Success", Response: response}
}

func GetErrorResponse(message string, response any) models.ApiResponse {
	return models.ApiResponse{Status: -1, Message: message, Response: response}
}

func ReturnJsonFromRows(rows *sql.Rows) (fJson []interface{}, fError *mssql.Error) {
	columnTypes, err := rows.ColumnTypes()

	if err != nil {
		return nil, &mssql.Error{Message: err.Error()}
	}

	count := len(columnTypes)
	finalRows := []interface{}{}

	for rows.Next() {

		scanArgs := make([]interface{}, count)

		for i, v := range columnTypes {

			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
				break
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
				break
			case "INT4":
				scanArgs[i] = new(sql.NullInt64)
				break
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := rows.Scan(scanArgs...)

		if err != nil {
			return nil, &mssql.Error{Message: err.Error()}
		}

		masterData := map[string]interface{}{}

		for i, v := range columnTypes {

			if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
				masterData[v.Name()] = z.Bool
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullString); ok {
				masterData[v.Name()] = z.String
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
				masterData[v.Name()] = z.Int64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
				masterData[v.Name()] = z.Float64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
				masterData[v.Name()] = z.Int32
				continue
			}

			masterData[v.Name()] = scanArgs[i]
		}

		finalRows = append(finalRows, masterData)
	}
	return finalRows, nil
	// z, err := json.Marshal(finalRows)
	// if err != nil {
	// 	return nil, &mssql.Error{Message: err.Error()}
	// } else {
	// 	return string(z), nil
	// }
	// return z, nil
}
