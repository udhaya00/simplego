package main
import "fmt"
func main() {
var x=[]int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}
var y int
var n=len(x)
for i:=0;i<n;i++{
for j:=0;j<n;j++{
if x[i]<x[j]{
	y=x[i]
	x[i]=x[j]
	x[j]=y
}
}
}
fmt.Printf("The Smallest number is : %d",x[0])
}


	
 	