package main
import (
"fmt"
"FACT"
"FIB"
)
func main() {
i:=1
var n int
for i!=0 {
fmt.Println("Enter an option : ")
fmt.Println("1.Factorial")
fmt.Println("2.Fibonacci series")
fmt.Println("3.exit")
fmt.Scanln(&n)
switch(n) {
case 1 : var a int
	 fmt.Println("Enter a number to find the factorial : ") 
	 fmt.Scan(&a)    
         fmt.Println(FACT.factorial(a))
	 break
case 2 : var b int
	 fmt.Println("Enter a number to find fibonacci series:  ")
         fmt.Scanln(&b)
	 for j:=0;j<=b;j++ {
	 fmt.Println(FIB.fib(j))
	 }
	 break 
case 3 : i=0
	 break
}
}
}