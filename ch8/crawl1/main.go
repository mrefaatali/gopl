package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

var tokens = make(chan struct{}, 100)

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
  tokens <- struct{}{} //acquire a token
	list, err := links.Extract(url)
  <-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main(){
  worklist := make(chan []string)
  var n int
  n++

  go func(){worklist <- os.Args[1:]}()

  seen := make(map[string]bool)
  for ;n>0; n--{
    list := <-worklist
    for _, link := range list {
      if !seen[link] {
        seen[link]=true
        n++
        go func(link string){
          worklist <- crawl(link)
        }(link)
      }
    }
  }

/*
  for list := range worklist {
    for _, link := range list{
      if !seen[link]{
        seen[link] = true
        go func(link string){
          worklist <- crawl(link)
        }(link)
      }
    }
  }
*/
}