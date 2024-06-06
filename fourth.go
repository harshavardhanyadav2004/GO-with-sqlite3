package show
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func DeleteTable(db *sql.DB) (int64 ,error) {
	fmt.Println("Deleting the user")
	id, err:= db.Exec("DELETE FROM Students WHERE  name = ?","Harsha") 
	if err!=nil{
		log.Fatal(err)
		return 0,err
	}
	new_id , err := id.RowsAffected()
	if  err!=nil {
		 log.Fatal(err)
		 return 0,err
	}
	return new_id , err

}

