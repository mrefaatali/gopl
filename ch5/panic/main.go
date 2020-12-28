package main

import (
	"fmt"
	"math/rand"
)

const (
	Spades = iota + 1
	Hearts
	Diamonds
	Clubs
	Joker
)

func main() {
	for i := 12; i > 0; i-- {
		switch n, s := drawCard(); s {
		case Spades:
			fmt.Println(n, "spades")
		case Hearts:
			fmt.Println(n, "hearts")
		case Diamonds:
			fmt.Println(n, "diamonds")
		case Clubs:
			fmt.Println(n, "clubs")
		default:
			panic(fmt.Sprintf("invalid suit %v", s))
		}
	}

}

func drawCard() (n int, s int) {
	n = rand.Intn(12) + 1
	s = rand.Intn(5) + 1
	return
}

func suit(_ int, s int) int {
	return s
}
