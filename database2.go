package main
 import (
 _ "github.com/go-sql-driver/mysql"
 "database/sql"
 "fmt"
 "log"
 )

 func main(){
 db, err := sql.Open("mysql", "rootadmin123@(localhost:8080)/UV")
 if err != nil {
 log.Fatal(err)
 }else{
 fmt.Println("Connection Established")
 }
 var (
 id int
 name string
 )
 rows,err:=db.Query("select id, username from user where id = ?", 1)
 if err!=nil{
 log.Fatal(err)
 }
 Defer rows.Close()
 for rows.Next(){
 err:=rows.Scan(&id,&name)
 if err != nil {
 log.Fatal(err)
 }
 fmt.Println(id, name)
 }
 err = rows.Err()
if err != nil {
 log.Fatal(err)
}
 defer db.Close()
 }