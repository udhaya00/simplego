package main
import "fmt"
func main() {
var x,y,a,b string
fmt.Println("Enter your first name")
fmt.Scanln(&x)
fmt.Println("Enter your last name")
fmt.Scanln(&y)
a,b=swap(x,y)
fmt.Println("After Swap: ",a,b)
}
func swap(x,y string)(string,string) {
var temp string
temp=x
x=y
y=temp
return x,y
}