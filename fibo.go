package main
import "fmt"
func fibonacci() chan int {
  c := make(chan int)
  go func() { 
    for i, j := 0, 1; ; i, j = i+j,i {
        c <- i
    }
  }()
  return c
}
func main() {
    c := fibonacci()
    for n := 0; n < 12 ; n++ {
        fmt.Printf("%d ", <- c)
    }
}