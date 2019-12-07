package main
import "fmt"
func main() {
var a[10] int
for i:=0;i<10;i++ {
a[i]=i+100
for j:=0;j<10;j++ {
fmt.Printf("element(%d)\n",j,a[j])
}
}
}