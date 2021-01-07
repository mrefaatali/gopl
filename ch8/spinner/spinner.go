package main

import (
	"fmt"
	"time"
)

func main() {
  go spinner(100 * time.Microsecond)
  const n = 45
  fibN := fib(n)
  fmt.Printf("\rFibbonaci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int64) int64 {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
