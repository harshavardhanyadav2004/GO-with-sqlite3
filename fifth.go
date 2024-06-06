package show 
import (
	"fmt"
	"log"
	"database/sql"
	 _ "github.com/mattn/go-sqlite3"

)
func SelectUser(db *sql.DB){
    rows ,err:=  db.Query("SELECT * FROM Students")
	if err!= nil{
		 log.Fatal(err)
		 return 
	}
	defer rows.Close()
	for rows.Next(){
		 var name string 
		 var email string 
		 var id int 
		 if err:=rows.Scan(&id, &email ,&name) ; err!=nil {
			 log.Fatal(err);
			 return 
		 }
		  fmt.Println(id,email,name) 

	}
	if  err:= rows.Err() ; err != nil {
		log.Fatal(err)
		return
	}
		
}