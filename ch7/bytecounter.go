package main

import "fmt"

//a type with a method that matches the signature of io.Write
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (n int, err error) {
	*c += ByteCounter(len(p)) // to convert int to name type ByteCounter
	defer func () {fmt.Printf("Write used and read %v: with length %d\n", p, n)}()
  n, err = len(p), nil
  return
}

/*
func main() {
	var c ByteCounter
	l,e := c.Write([]byte("hello"))
	fmt.Println(l,e,c)

  //c=0
  var name = "Dolly"
  fmt.Fprintf(&c, "hello, %s", name)
  fmt.Println(c)
}

func (c ByteCounter) String() string{
  return fmt.Sprintf("%d bytes counted", c)
}
*/