package main

import (
)

func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    printVar("a[0]", a[0])
    printVar("a[1]", a[1])
    printVar("a", a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    printVar("primes", primes)
}
