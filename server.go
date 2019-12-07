package main
import (
"fmt"
"net"
)
func main() {
listener,_:=net.Listen("tcp",":8081")
for {
conn,_:=listener.Accept()
go doServerStuff(conn)
}
}
func doServerStuff(conn net.Conn){
for {
buf:=make([]byte,512)
_,err:=conn.Read(buf)
if err!=nil {
fmt.Println("Error Reading",err.Error())
return
}
fmt.Printf("Recieved data : %v",string(buf))
}
}