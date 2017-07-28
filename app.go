package main

import (
    "github.com/dwitp/golang/src/u"
)

func main() {
    s := make([]int, 5)
    u.D("s")
    u.S(s)

    b := make([]int, 0, 5)
    u.D("b")
    u.S(b)

    c := b[:2]
    u.D("c")
    u.S(c)

    d := c[2:5]
    u.D("d")
    u.S(d)
}
