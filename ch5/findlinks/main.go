package main

import (
	"fmt"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findlinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks: %v", err)
			continue
		}
		fmt.Println("\n###", url, "###")
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
