package geometry

import(
  "math"
  "image/color"
) 

type Point struct{ X, Y float64 }

//traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Point method
func (p Point) Distance(q Point) float64 {
  return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64){
  p.X *= factor
  p.Y *= factor
}

type ColoredPoint struct {
  Point
  Color color.RGBA
}
