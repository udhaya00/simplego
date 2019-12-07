package main
import "fmt"
type point struct {
x int
y int
}
func (p point)Abs() int {
	return p.x*p.x+p.y*p.y
}
func main() {
p:=new(point)
p.x=10
p.y=5
fmt.Println("The result is : ",p.Abs())
}