package tables

import "udit/api-padhai/models"

var todoTables []models.Tables

func populateTablesStructure() {
	todoTables = append(todoTables, models.Tables{TableName: "colors", TableCreationQuery: "CREATE TABLE [dbo].[colors]([color_id] [int] IDENTITY(1,1) NOT NULL,[color_name] [varchar](50) NULL,[color_value] [varchar](50) NULL,[is_light] [bit] NULL,[create_id] [int] NULL,[created_on] [datetime] default getdate())"})
}

func GetTables() []models.Tables {
	if len(todoTables) <= 0 {
		populateTablesStructure()
	}
	return todoTables
}
