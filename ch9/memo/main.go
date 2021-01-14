package main

import (
	"fmt"
	"io/ioutil"
	"log"
	memo "main/ch9/memo/memo3"
	"net/http"
	"os"
	"sync"
	"time"
)

//a function that we want to avoid calling repeatidly
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLS() []string {
	return os.Args[1:]
}

func main() {
	m := memo.New(httpGetBody)
	var n sync.WaitGroup

	for _, url := range incomingURLS() {
		n.Add(1)
		time.Sleep(1 * time.Second)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

/*
go run -race -v ./ch9/memo https://golang.org https://godoc.org https://play.golang.org https://golang.org https://godoc.org https://play.golang.org https://golang.org https://godoc.org https://play.golang.org
*/
