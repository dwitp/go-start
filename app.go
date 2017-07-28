package main

import (
    "github.com/dwitp/golang/src/u"
)

func main() {
    q := []int{2, 3, 5, 7, 11, 13}
    u.D("q", q)

    r := []bool{true, false, true, true, false, true}
    u.D("r", r)

    s := []struct{
        i int
        b bool
    }{
        {2, true},
        {3, false},
        {5, true},
        {7, true},
        {11, false},
        {13, true},
    }
    u.D("s", s)
}
