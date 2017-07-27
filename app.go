package main

import (
)

type Vertex struct {
    X int
    Y int
}

func main() {
    vertex := Vertex{1, 2}
    printVar("vertex", vertex)
}
