package main

import (
    "github.com/dwitp/go-start/src/config"
)

func main() {
    i, j := 42, 2701

    p := &i
    printVar("p", p)
    printVar("*p", *p)

    *p = 21
    printVar("i", i)

    p = &j
    *p = *p / 37
    printVar("j", j)

    v := Vertex{3, 4}
    printVar("v", v)
    printVar("v.Abs()", v.Abs())
}
