package main 
import (
	 "log"
	 "fmt"
	 "database/sql"
   _ "github.com/mattn/go-sqlite3"
	 "show"

)
const file string = "Another_database.db"

func main() {
	fmt.Println("Connecting into the database")
	new_db,err:=sql.Open("sqlite3",file)
	if err!=nil{
		log.Fatal(err)
		return 
	}
	show.CreateTable(new_db)
	new_id , err:= show.InsertIntoTable(new_db,1,"Harsha10012004@gmail.com","Harsha")
	if err!=nil {
		log.Fatal(err)
		return 
	}
	fmt.Println("The id is ",new_id)
	another_id,err := show.InsertIntoTable(new_db,2,"vardhan102@gmail.com","Vardhan")
	if err!=nil {
		log.Fatal(err)
		return
	}
	fmt.Println("The new_id is ",another_id)
	other_id , err :=  show.UpdateId(new_db)
	if err!=nil {
		log.Fatal(err)
		return 
	}
	fmt.Println("The update id is ",other_id)
	new_id_1 , err := show.DeleteTable(new_db)
	if err!=nil {
		  log.Fatal(err)
		  return
	}
	fmt.Println("The rows Affected is ",new_id_1)
	show.SelectUser(new_db)
	fmt.Println("Rows Selected Succesfully")
}
