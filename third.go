package show
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
func UpdateId(db *sql.DB)(int64,error){
      fmt.Println("Connecting into the database")
	  fmt.Println("Updating the table in the database")
	  result ,err:= db.Exec("UPDATE Students SET email = ? WHERE name = ?","harsha10012004@gmail.com","Harsha ")
	  if err!=nil {
		return 0,err 
	  }
	  no_of_rows,err := result.RowsAffected()
	  if err!=nil {
		 log.Fatal(err)
		 return 0,err
	  }
	  fmt.Println("The Rows effected aree")
	  return no_of_rows,nil

}