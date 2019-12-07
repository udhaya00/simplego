package main
import "fmt"
func factorial(x int) int {
	return x * factorial(x-1)
}
func main() {
	x := int(5)
	calcFactorial := factorial(x)
	fmt.Println(calcFactorial)
}