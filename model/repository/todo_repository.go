package repository

import (
	"fmt"
	baseDb "golag-study/model"
	"log"
)

func Index() {
	db, err := baseDb.New()

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(db)
	// rows, err := db.
	// 	Query("SELECT id, title, content FROM todo ORDER BY id DESC")
}