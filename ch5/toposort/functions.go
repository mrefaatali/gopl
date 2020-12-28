package main
/*
import (
  "fmt"
)

func main() {
	f := square
	fmt.Printf("f: %T \n", f) //func(int) int
	fmt.Println(f(3))

	var g func(int) int
	fmt.Printf("g: %T\n", g)
	if g != nil {
		g(1)
	}

  h := squares()
  fmt.Printf("h: %T\n", h)
  fmt.Println(h())
  fmt.Println(h())
  fmt.Println(h())
}
*/
func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
func squares() func() int {
  var x int
  return func() int {
    x++
    return x*x
  }
}