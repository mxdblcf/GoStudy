package main
import (
	"database/sql"
 "fmt"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	db, err := sql.Open("mysql", "mxd:@#MA847547125**@(127.0.0.1:3306)/mxd")
db.Ping()
if err!=nil{
	fmt.Printf("err")
}
defer db.Close()
	query, _ := db.Query("select * from user")
var id string
	var name string
	var a1 int
	for query.Next(){
		  query.Scan(&id,&name,&a1)
		fmt.Println(id,name,a1)

	}
}
