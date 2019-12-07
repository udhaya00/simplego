package main
import "fmt"
func main() {
var a,b,c float64
fmt.Println("Enter three heights")
fmt.Scanln(&a)
fmt.Scanln(&b)
fmt.Scanln(&c)
fmt.Println("The second tallest is: ",highest(a,b,c))
}
func highest(a,b,c float64)float64 {
if a>b && a<c{
	return a
}else
if b>a && b<c{
	return b
}else{
	return c
}
}