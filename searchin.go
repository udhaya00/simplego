package main
import(
 "fmt"
 "strconv"
 "regexp"
)
func main() {
searchin := "John:2578.34 William:4567.23 Steve:5632.18"
pat:="(0-9)+.(0-9)+"
f :=func (s string)string {
v,_:=strconv.ParseFloat(s,32)
     return strconv.FormatFloat(v*2,'f',2,32)
}
if ok,_:=regexp.Match(pat,[]byte(searchin));ok {
fmt.Println("Match Found")
}
re,_:=regexp.Compile(pat)
str :=re.ReplaceAllString(searchin,"##.#")
fmt.Println(str)
str2:=re.ReplaceAllStringFunc(searchin,f)
fmt.Println(str2)
}