package main

import (
	"fmt"
	db "golag-study/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("successfully migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(
		&db.Task{},
	)
}
