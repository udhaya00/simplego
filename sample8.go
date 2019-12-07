package main
import(
"fmt"
"strings"
)
func main(){
fmt.Println("Contains:",strings.Contains("Programming","gram"))
fmt.Println("Count:",strings.Count("Programming","m"))
fmt.Println("HasPrefix:",strings.HasPrefix("Programming","Pr"))
fmt.Println("HasSuffix:",strings.HasSuffix("Programming","ing"))
fmt.Println("Index:",strings.Index("Programming","g"))
fmt.Println("Join:",strings.Join([]string{"Go","Programming"},"-"))
fmt.Println("Repeat:",strings.Repeat("Go",5))
fmt.Println("Replace:",strings.Replace("Chello","l","n",-1))
fmt.Println("Replace:",strings.Replace("Chello","l","n",1))
fmt.Println("Split:",strings.Split("172.16.16.200","."))
fmt.Println("ToLower:",strings.ToLower("PROGRAMMING"))
fmt.Println("ToUpper:",strings.ToUpper("programming"))
fmt.Println()
fmt.Println("Len:",len("Programming"))
fmt.Println("Char:","Programming"[5])
}