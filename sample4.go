package main
import(
"fmt"
"strings"
)
func main(){
fmt.Println(strings.Compare("Hello","Hello"))
fmt.Println(strings.Compare("Hello","Cello"))
fmt.Println(strings.Compare("Cello","Hello"))
fmt.Println(strings.Compare("Hello","hello"))
}