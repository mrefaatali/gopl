package main

import (
	"fmt"
)

func main() {
	f := square
	fmt.Printf("f: %T \n", f) //func(int) int
	fmt.Println(f(3))

	var g func(int) int
	fmt.Printf("g: %T", g)
	if g != nil {
		g(1)
	}
}

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
