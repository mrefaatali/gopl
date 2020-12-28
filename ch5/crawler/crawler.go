package main

import (
	"fmt"
	"log"
	"main/ch5/crawler/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
  return list
}
