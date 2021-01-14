package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, n *sync.WaitGroup,fileSizes chan<- int64) {
  defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
      n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n,fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 50)

func dirents(dir string) []os.FileInfo {
  sema <- struct{}{} //acquire a tokken
  defer func(){<-sema}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.2f KB\n", nfiles, float64(nbytes)/1e3)
}

var verbose = flag.Bool("v", false, "show verbose progress message")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var n sync.WaitGroup
  for _, root := range roots{
    n.Add(1)
    go walkDir(root, &n, fileSizes)
  }
  go func(){
    n.Wait()
    close(fileSizes)
  }()

	var tick <-chan time.Time

	if *verbose {
		tick = time.Tick(time.Millisecond)
	}

	var nfiles, nbytes int64
M:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break M
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}
