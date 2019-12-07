package main
import(
"net"
"bufio"
"fmt"
"os"
"strings"
)
func main() {
conn,err:= net.Dial("tcp",":8081")
inputReader:=bufio.NewReader(os.Stdin)
fmt.Println("What is your name")
clientname,_:=inputReader.ReadString('\n')
trimmedClient:=strings.Trim(clientname,"\r\n")
for{
fmt.Println("What to send to server?")
input,_:=inputReader.ReadString('\n')
trimmedinput:=strings.Trim(input,"\r\n")
if trimmedinput=="Q"{
return 
}
if err!=nil {
fmt.Println("Error Reading",err.Error())
return
}
_,err=conn.Write([]byte(trimmedClient+" says : "+trimmedinput))
}
}