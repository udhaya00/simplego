package main
import "fmt"
import "time"
func factorial(n int) {
	var Val int = 1      
        for i:=1; i<=n; i++ {
            Val *= i
        }         
   fmt.Println("Factorial is:  ",Val) 
}
func main(){
	var n int
	fmt.Print("Enter a number to find the factorial : ") 
	fmt.Scan(&n)    
        go factorial(n)
	time.Sleep(1e9)
}