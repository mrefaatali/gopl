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

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 50)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: //acquire a tokken
	case <-done:
		return nil
	}
  //sema <- struct{}{}
	defer func() { <-sema }()
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

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time

	if *verbose {
		tick = time.Tick(1000 * time.Microsecond)
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
    //done <- struct{}{}
	}()

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
      time.Sleep(10 * time.Millisecond)
			printDiskUsage(nfiles, nbytes)
		case <-done:
			//drain file sizes
			for range fileSizes {
				//do nothing
			}
			panic("TESTING done")
		}
	}
	printDiskUsage(nfiles, nbytes)
}
