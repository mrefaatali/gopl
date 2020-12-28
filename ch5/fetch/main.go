package main

import (
  "os"
  "log"
  "fmt"
)

func main() {
	for _, url := range os.Args[1:] {
    fn, n, err := fetch(url)
    if err !=nil {
      log.Println(err)
    }
    fmt.Println(fn)
    fmt.Println(n)
	}
}
