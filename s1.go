package main
import (
"fmt"
"SP"
)
func main() {
var n float64
var a int
var b string
var s SP.student
	
fmt.Println("Enter the details of student ")
fmt.Println("Enter the Rollno: ")
fmt.Scanln(&s.rollno)
fmt.Println("Enter the CGPA: ")
fmt.Scanln(&s.cgpa)
fmt.Println("Enter the Name: ")
fmt.Scanln(&s.name)

fmt.Println("After sorting...")

fmt.Println("Details of student ")
fmt.Println("Name   : ",s.name)
fmt.Println("CGPA   : ",s.cgpa)
fmt.Println("RollNo : ",s.rollno)
}
