package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
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
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var tick <-chan time.Time

	if *verbose {
		tick = time.Tick( time.Millisecond)
	}

	var nfiles, nbytes int64
M:
  for{
    select{
      case size, ok := <-fileSizes:
        if !ok{
          break M
        }
        nfiles ++
        nbytes += size
      case <-tick:
        printDiskUsage(nfiles, nbytes)
    }
  }
  printDiskUsage(nfiles, nbytes)
}
