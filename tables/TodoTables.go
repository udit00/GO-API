package tables

import "udit/api-padhai/models"

var todoTables []models.Tables

func populateTablesStructure() {
	todoTables = append(todoTables, table_User())
	todoTables = append(todoTables, table_Todo())
	todoTables = append(todoTables, table_TodoType())
	todoTables = append(todoTables, table_TodoStatus())
	todoTables = append(todoTables, table_Colors())
	todoTables = append(todoTables, table_LoginLogs())
}

func table_User() models.Tables {
	alterQueries := []string{
		"ALTER TABLE [dbo].[users] ADD PRIMARY KEY CLUSTERED ( [user_id] ASC )",
		"ALTER TABLE [dbo].[users] ADD  DEFAULT (NULL) FOR [display_picture]",
		"ALTER TABLE [dbo].[users] ADD  DEFAULT (getdate()) FOR [created_on]",
		"ALTER TABLE [dbo].[users] ADD  DEFAULT (NULL) FOR [email_id]",
		"ALTER TABLE [dbo].[users] ADD  DEFAULT (NULL) FOR [mobile_no]",
		"ALTER TABLE [dbo].[users] ADD  DEFAULT ((1)) FOR [is_active]",
		"ALTER TABLE [dbo].[users] ADD  DEFAULT ((0)) FOR [is_premium]",
	}
	table := models.Tables{
		TableName: "users",
		TableCreationQuery: "CREATE TABLE [dbo].[users]( " +
			"[user_id] [int] IDENTITY(1,1) NOT NULL, " +
			"[name] [varchar](50) NOT NULL, " +
			"[pass] [varchar](100) NOT NULL, " +
			"[display_picture] [varchar](500) NULL, " +
			"[created_on] [datetime] NOT NULL, " +
			"[firebase_token] [text] NULL, " +
			"[email_id] [varchar](50) NULL, " +
			"[mobile_no] [varchar](20) NULL, " +
			"[is_active] [bit] NULL, " +
			"[is_premium] [bit] NOT NULL )",
		AlterTableQueries: alterQueries[:],
	}
	return table
}

func table_Todo() models.Tables {
	alterQueries := []string{
		"ALTER TABLE [dbo].[todo] ADD PRIMARY KEY CLUSTERED ( [user_id] ASC, [todo_id] ASC )",
		"ALTER TABLE [dbo].[todo] ADD  DEFAULT (getdate()) FOR [created_on]",
		"ALTER TABLE [dbo].[todo] ADD  DEFAULT (getdate()+(1)) FOR [target]",
		"ALTER TABLE [dbo].[todo]  WITH CHECK ADD FOREIGN KEY([user_id]) REFERENCES [dbo].[users] ([user_id])",
	}
	table := models.Tables{
		TableName: "todo",
		TableCreationQuery: "CREATE TABLE [dbo].[todo]( " +
			"[todo_id] [int] NOT NULL, " +
			"[title] [varchar](100) NOT NULL, " +
			"[description] [text] NOT NULL, " +
			"[user_id] [int] NOT NULL, " +
			"[created_on] [datetime] NOT NULL, " +
			"[type_id] [int] NOT NULL, " +
			"[target] [datetime] NULL,  " +
			"[completion_status_id] [int] NULL ) ",
		AlterTableQueries: alterQueries,
	}
	return table
}

func table_TodoType() models.Tables {
	alterQueries := []string{
		"ALTER TABLE [dbo].[todo_type] ADD PRIMARY KEY CLUSTERED ( [create_id] ASC, [type_id] ASC )",
		"ALTER TABLE [dbo].[todo_type] ADD  DEFAULT (getdate()) FOR [created_on]",
	}
	table := models.Tables{
		TableName: "todo_type",
		TableCreationQuery: "CREATE TABLE [dbo].[todo_type]( " +
			"[type_id] [int] NOT NULL, " +
			"[type_name] [varchar](100) NOT NULL, " +
			"[create_id] [int] NOT NULL, " +
			"[created_on] [datetime] NOT NULL, " +
			"[color_id] [int] NOT NULL )",
		AlterTableQueries: alterQueries,
	}
	return table
}

func table_TodoStatus() models.Tables {
	alterQueries := []string{
		"ALTER TABLE [dbo].[todo_status] ADD  DEFAULT (getdate()) FOR [created_on]",
	}
	table := models.Tables{
		TableName: "todo_status",
		TableCreationQuery: "CREATE TABLE [dbo].[todo_status]( " +
			"[status_id] [int] IDENTITY(1,1) NOT NULL, " +
			"[status_name] [varchar](20) NOT NULL, " +
			"[color_id] [int] NULL, " +
			"[create_id] [int] NOT NULL, " +
			"[created_on] [datetime] NOT NULL )",
		AlterTableQueries: alterQueries,
	}
	return table
}

func table_Colors() models.Tables {
	return models.Tables{
		TableName:          "colors",
		TableCreationQuery: "CREATE TABLE [dbo].[colors]([color_id] [int] IDENTITY(1,1) NOT NULL,[color_name] [varchar](50) NULL,[color_value] [varchar](50) NULL,[is_light] [bit] NULL,[create_id] [int] NULL,[created_on] [datetime] default getdate())",
		AlterTableQueries:  nil,
	}
}

func table_LoginLogs() models.Tables {
	var alterQueries []string
	alterQueries = append(alterQueries, "ALTER TABLE [dbo].[login_logs] ADD  DEFAULT (NULL) FOR [user_id]")
	alterQueries = append(alterQueries, "ALTER TABLE [dbo].[login_logs] ADD  DEFAULT (NULL) FOR [ip_address]")
	alterQueries = append(alterQueries, "ALTER TABLE [dbo].[login_logs] ADD  DEFAULT (getdate()) FOR [login_time]")
	alterQueries = append(alterQueries, "ALTER TABLE [dbo].[login_logs] WITH CHECK ADD FOREIGN KEY([user_id]) REFERENCES [dbo].[users] ([user_id])")
	table := models.Tables{
		TableName:          "login_logs",
		TableCreationQuery: "CREATE TABLE [dbo].[login_logs]([user_id] [int] NULL,[ip_address] [varchar](100) NULL,[login_time] [datetime] NULL,[session_id] [varchar](100) NOT NULL,[platform] [text] NULL)",
		AlterTableQueries:  alterQueries,
	}
	return table
}

func GetTables() []models.Tables {
	if len(todoTables) <= 0 {
		populateTablesStructure()
	}
	return todoTables
}
