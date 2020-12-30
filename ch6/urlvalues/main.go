package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["items"])

  //m = nil
  fmt.Println(m.Get("item"))
  v, check := m["test"]
  fmt.Println(v,check,m)
  m["items"] = []string{"11", "22"}
  fmt.Println(m)
  m.Add("items", "3")

  a := map[int]int{1: 11, 2: 22}
  iv, icheck := a[3]
  fmt.Println(iv, icheck, a)
}