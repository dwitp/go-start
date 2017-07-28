package main

import (
    "github.com/dwitp/golang/src/u"
)

func main() {
    var s[]int
    u.D(s, len(s), cap(s))
    if s == nil {
        u.D("nil!")
    }
}
