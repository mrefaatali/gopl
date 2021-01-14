package main

import (
	"fmt"
	"time"
	//"os"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go Counter(naturals, 10)

	//squarer
	go Squarer(naturals, squares)

	//printer
	Printer(squares)

}

func Counter(pub chan<- int, n int) {
	fmt.Println("Counter called")
	for x := 0;x<n ; x++ {
    pub <- x
    time.Sleep(100*time.Millisecond)
	}
  fmt.Println("naturals closed")
  close(pub)
}

func Squarer(sub <-chan int, pub chan<- int) {
	fmt.Println("Squarer called")
	for x := range(sub){
		pub <- x * x
	}
  fmt.Println("Squarer closed")
	close(pub)
}

func Printer(sub <-chan int) {
	fmt.Println("Printer called")
	for x := range(sub){
		fmt.Println("Printer: ",x)
	}
  fmt.Println("Printer finished")
}
