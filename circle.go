package main
import "fmt"
func main() {
var a,x,y float64
fmt.Println("Enter the radius of the circle")
fmt.Scanln(&a)
x,y=circle(a)
fmt.Printf("Area= %f\nPerimeter= %f",x,y)
}
func circle(a float64)(float64,float64){
x:=3.14*a*a
y:=2*3.14*a
return x,y
}

