package main

import (
    "github.com/dwitp/golang/src/u"
)

func fibonacci() func() int {
    cur, next := 1, 0
    return func() int {
        cur, next = next, cur+next
        return cur
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        u.D(f())
    }
}
