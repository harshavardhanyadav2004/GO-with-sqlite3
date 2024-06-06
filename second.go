package show
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
func InsertIntoTable(db *sql.DB,inc int ,email string , name string) (int64 , error) {
	fmt.Println("Inserting into table into the database")
	fmt.Println("Inserting the records in the database")
	result , err:= db.Exec("INSERT INTO Students (id,email ,name) VALUES (?,?,?)",inc,email,name)
	if err!=nil {
		log.Fatal(err)
		return 0,err
	}
	id ,err:=result.LastInsertId()
	if err!=nil {
		return 0,err
	}
	return  id, err
	
}
