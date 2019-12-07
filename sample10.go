package main
import(
"fmt"
"strings"
)
func main(){
fmt.Println("Enter a IP adress")
var a string
fmt.Scan(&a)
var b=strings.Split(a,".")
var c=b[0]
if c>"1" && c<"124"{
fmt.Println("Class A")
}
if c>"124" && c<"136"{
fmt.Println("Class B")
}
if c>"136" && c<"192"{
fmt.Println("Class C")
}
if c>"192" && c<"256"{
fmt.Println("Class D")
}
}