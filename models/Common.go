package models

import (
	"awesomeProject/database"
	"database/sql"
)

type Common struct {
	Db        *sql.DB
	TableName string
}

func (c *Common) FindAll(columns []string) *sql.Rows {
	rows, _ := c.Db.Query("select * from " + c.TableName)
	return rows
}

func DatabaseConnectionInit(tableName string) *Common {
	dbConnect := database.DbConnect()
	return &Common{Db: dbConnect, TableName: tableName}
}
