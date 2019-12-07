package main
import "fmt"
func main() {
var n,a,b,c int
fmt.Println("Enter the option")
fmt.Println("1.Add\n2.Subtract\n3.Multiply\n4.Divide")
fmt.Scanln(&n)
switch n {
case 1: fmt.Println("Enter the the First value")
	fmt.Scanln(&a)
	fmt.Println("Enter the second value")
	fmt.Scanln(&b)
	Add(a,b)
	break
case 2: Sub()
	break
case 3: fmt.Println("Enter the the First value")
	fmt.Scanln(&a)
	fmt.Println("Enter the second value")
	fmt.Scanln(&b)
	c=Mul(a,b)
	fmt.Println("The Answer is: ",c)
	break 
case 4: fmt.Println("The Answer is: ",Div())
	break
default: fmt.Println("invalid case")
	break
}
}
func Add(x int,y int) {
	fmt.Println("The Answer is: ",x+y)
}
func Sub() {
	var x,y int
	fmt.Println("Enter the the First value")
	fmt.Scanln(&x)
	fmt.Println("Enter the second value")
	fmt.Scanln(&y)
	fmt.Println("The Answer is: ",x-y)
}
func Mul(a int,b int)int {
	return a*b
}
func Div()int {
	var x,y int
	fmt.Println("Enter the the First value")
	fmt.Scanln(&x)
	fmt.Println("Enter the second value")
	fmt.Scanln(&y)
	return x/y
}