package handlers

import (
	"fmt"
	"log"
	"submission-godb/database"
)

func isExists(tableName string, columnName string, value interface{}) bool {
	db := database.ConnectDb()
	defer db.Close()
	var exists bool
	sqlStatement := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s=$1)", tableName, columnName)
	err := db.QueryRow(sqlStatement, value).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	return exists
}
