package main
import "fmt"
func main() {
var a=[]int{4,6,3,12,10}
reverse(a[])
}
func reverse(a []int) {
b:=len(a)
for i:=b-1;i>=0;i--{
fmt.Printf("%d,"a[i])
}
}
