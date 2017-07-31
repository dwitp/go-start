package main

import (
    "github.com/dwitp/golang/src/u"
)

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        u.D(pos(i), neg(-2*i))
    }
}
