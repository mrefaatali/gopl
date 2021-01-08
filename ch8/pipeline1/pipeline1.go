package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go Counter(naturals)

	//squarer
	go Squarer(naturals, squares)

	//printer
	Printer(squares)

}

func Counter(send chan int) {
  fmt.Println("Counter called")
	for x := 0; ; x++ {
		send <- x
	}
}

func Squarer(recv, send chan int) {
	fmt.Println("Squarer called")
  for {
		x := <-recv
		send <- x
	}
}

func Printer(recv chan int) {
	fmt.Println("Printer called")
  for {
		fmt.Println(<-recv)
	}
}
