package main

import (
    "github.com/dwitp/golang/src/u"
    "math"
)

func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    u.D(hypot(5, 12))
    u.D(compute(hypot))
    u.D(compute(math.Pow))
}
