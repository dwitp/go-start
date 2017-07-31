package main

import (
    "github.com/dwitp/golang/src/u"
)

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex {
        40.68433, -74.39967,
    }
    u.D(m["Bell Labs"])
}
