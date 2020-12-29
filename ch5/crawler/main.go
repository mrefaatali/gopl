package main

import (
	"fmt"
	"os"
)

/*
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
*/

//crawler
/*
func main(){
  breadthFirst(crawl, os.Args[1:])
}
*/

func main() {
	for _, url := range os.Args[1:] {
    if title(url) != nil {
			fmt.Fprintf(os.Stderr, "title in %s : %v", url, title(url))
		}
	}
}
