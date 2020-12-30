package main

import (
	g "main/ch6/geometry"
	"fmt"
  "image/color"
)

func main() {
	p := g.Point{1, 2}
	q := g.Point{4, 6}
  p1 := g.Point{0,0}
  p2 := g.Point{5,5}

	fmt.Println(g.Distance(p, q))
	fmt.Println(p.Distance(q))

  path := g.Path{p1, p2, p,q}
  fmt.Println(path.Distance())

  path2 := g.Path{
    {1,1},
    {5,1},
    {5,4},
    {1,1},
  }
  fmt.Println(path2.Distance())

  r := &p
  r.ScaleBy(3)
  fmt.Printf("%T %T\n", p, p.ScaleBy)
  fmt.Println(r)
  fmt.Println(*r)

  var cp g.ColoredPoint
  cp.X = 1
  fmt.Println(cp)
  cp.ScaleBy(3)
  fmt.Println(cp)
  red := color.RGBA{255,0,0,255}
  blue := color.RGBA{0,0,255,255}
  green := color.RGBA{0,255,0,255}
  cp.Color = green
  fmt.Println(cp)
  var  pc = g.ColoredPoint{p, red}
  var qc = g.ColoredPoint{q, blue}
  fmt.Println(pc)
  fmt.Println(qc)
  fmt.Println(pc.Distance(qc.Point))
}
