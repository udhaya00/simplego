package main
import "fmt"
func main() {
var n,a int
fmt.Println("Enter the No. of Units")
fmt.Scanln(&n)
if n<100{
firstrange(n)
}
if n>100 && n<300{
a=n-100
fmt.Println(firstrange(100)+secondrange(a))
}
if n>300 && n<500{
a=n-300
fmt.Println(firstrange(100)+secondrange(200)+thirdrange(a))
}
if n>500{
a=n-500
fmt.Println(firstrange(100)+secondrange(200)+thirdrange(200)+fourthrange(a))
}
}
func firstrange(n int)int {
	return n*3
}
func secondrange(a int)int {
	return a*5
}
func thirdrange(a int)int {
	return a*8
}
func fourthrange(a int)int {
	return a*10
}
	