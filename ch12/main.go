package main

import (
	"fmt"
	"reflect"
)

func main() {
  type mohamed struct{
    children int
  }
  i := mohamed{2}
 func (m Mohamed) ay7aga() int {return m.children}
/*
  go func(){
    i <- 1
    close(i)
  }()
*/

	t := reflect.TypeOf(i)
  v := reflect.ValueOf(i)
  m := t.NumMethod()
  in := v.Interface()

  printer(t)
  printer(v)
  printer(m)
  printer(in)

/*
  go func(){
    dump := <-i
    printer(dump)
    }()
*/
}

func printer(v interface{}){
  fmt.Printf("-value: %8v\ttype: %T\n", v,v)
}
