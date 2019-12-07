package main
import "fmt"
func main() {
var n,a,b int
fmt.Println("Enter a number")
fmt.Scanln(&n)
a,b=eq(n)
fmt.Printf("Y= %d, Z= %d",a,b)
}
func eq(n int)(int,int) {
y:=(n*n)+n+5
z:=(n*n*n)+(n*2)+11
return y,z
}
