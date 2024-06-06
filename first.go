package show

import (
	 "database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"

)
func CreateTable(db *sql.DB){
	const create string = `
	CREATE TABLE Students (
		id INTEGER 	NOT NULL,
		email TEXT ,
		name TEXT 
	)
	`
	fmt.Println("Connecting to the database in the sqlite3")
	if _,err:= db.Exec(create) ; err!=nil {
		log.Fatal(err)
		return 
	}
	fmt.Println("Created a table in the database succesfully")
}