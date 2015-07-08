package main

import (
  "fmt"
  "math"
)

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}
//structs
type Circle struct {
  x float64
  y float64
  r float64
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

//methods
func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func (c *Circle) perimeter() float64 {
  return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l*(2) + w*(2)
}

//interfaces
type Shape interface {
  area() float64
  perimeter() float64
}

func main()  {
  r := Rectangle{0, 0, 10, 10}
  fmt.Println(r.area())
  fmt.Println(r.perimeter())

  c := Circle{0, 0, 5}
  fmt.Printf("%.2f\n", c.area())
  fmt.Printf("%.2f\n", c.perimeter())
}
