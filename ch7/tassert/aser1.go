package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
  fmt.Printf("%v of type %[1]T\n",w)
	f := w.(*os.File)
  fmt.Printf("%v of type %[1]T\n",f)
	rw := w.(io.ReadWriter)
	fmt.Printf("%v of type %[1]T\n",rw)
  c, ok := w.(*bytes.Buffer)
  fmt.Printf("%v of type %[1]T\n",ok)
  fmt.Printf("%v of type %[1]T\n",c)
}
