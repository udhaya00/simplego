package main
import (
"fmt"
"SP"
)
func main() {
var n float64
var a int
var b string
s:=[5]Student{}
for i:=0;i<5;i++ {	
fmt.Println("Enter the details of student ",i+1)
fmt.Println("Enter the Rollno: ")
fmt.Scanln(&s[i].Rollno)
fmt.Println("Enter the CGPA: ")
fmt.Scanln(&s[i].Cgpa)
fmt.Println("Enter the Name: ")
fmt.Scanln(&s[i].Name)
}
for x:=0;x<5;x++{
for y:=0;y<5;y++{
if s[x].Cgpa<s[y].Cgpa{
	n=s[x].Cgpa
	s[x].Cgpa=s[y].Cgpa
	s[y].Cgpa=n
	a=s[x].Rollno
	s[x].Rollno=s[y].Rollno
	s[y].Rollno=a
	b=s[x].Name
	s[x].Name=s[y].Name
	s[y].Name=b
}
}
}
fmt.Println("After sorting...")
for i:=0;i<5;i++ {
fmt.Println("Details of student ",i+1)
fmt.Println("Name   : ",s[i].name)
fmt.Println("CGPA   : ",s[i].cgpa)
fmt.Println("RollNo : ",s[i].rollno)
}
}