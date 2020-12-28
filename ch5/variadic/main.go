package main

import "fmt"

func main() {
	fmt.Printf("sum of 1,2,3 is: %d\n", sum(1, 2, 3))

  values := []int{1,2,3,4,5}
  fmt.Println(sum(values...))
}


