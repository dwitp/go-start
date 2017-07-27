package main

import (
)

type Vertex struct {
    X int
    Y int
}

var (
    v1 = Vertex{1, 2}   // has type Vertex
    v2 = Vertex{X: 1}   // Y:0 is implicit
    v3 = Vertex{}       // X:0 and Y:0
    p  = &Vertex{1, 2}  // has type *Vertex
)

func main() {
    printVar("v1", v1)
    printVar("p", p)
    printVar("v2", v2)
    printVar("v3", v3)
}
