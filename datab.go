package main
import(
"database/sql"
"log"
_"github.com/go-sql-driver/mysql"
)
func main() {
db,err :=sql.Open("mysql","rootadmin123@(localhost:3037)/UV")
if err!=nil {
log.Fatal(err)
}
if err :=db.Ping();err!=nil {
log.Fatal(err)
}
insert,err :=db.Query("INSERT INTO emp VALUES {'kumar',19,7}")
defer insert.Close()
}