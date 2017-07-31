package main

import (
    "github.com/dwitp/golang/src/u"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}
    u.D(v.Abs())
    u.D(AbsFunc(v))

    p := &Vertex{4, 3}
    u.D(p.Abs())
    u.D(AbsFunc(*p))
}
