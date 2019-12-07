package main
import(
"fmt"
"sort"
)
func main() {
var a[5]int
fmt.Println("Enter the values one by one")
for i:=0;i<5;i++{
fmt.Scanln(&a[i])
}
sort.Ints(a)
fmt.Println(a)
}