package main
import(
"fmt"
"sort"
) 	
func main() {
var x=[]int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}
sort.Ints(x)
fmt.Printf("The Smallest number is: %d",x[0])
}