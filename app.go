package main

import (
    "github.com/dwitp/golang/src/u"
)

func main() {
    s := []int{2, 3, 5, 7, 11, 13}

    s = s[1:4]
    u.D("s", s)

    s = s[:2]
    u.D("s", s)

    s = s[1:]
    u.D("s", s)
}
