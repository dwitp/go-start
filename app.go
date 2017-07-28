package main

import (
    "github.com/dwitp/golang/src/u"
)

func main() {
    var s[]int
    u.S(s)

    // append works on nil slices.
    s = append(s, 0)
    u.S(s)

    // The slice grows as needed.
    s = append(s, 1)
    u.S(s)

    // We can add more than one element at a time.
    s = append(s, 2, 3, 4)
    u.S(s)
}
