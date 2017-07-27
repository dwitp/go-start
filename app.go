package main

import (
    "fmt"
)

func main() {
    fmt.Println("~~~", "app.go main")
    // ==================================================
    i, j := 42, 2701

    p := &i
    printVar("p", p)
    printVar("*p", *p)

    *p = 21
    printVar("i", i)

    p = &j
    *p = *p / 37
    printVar("j", j)
}

func printVar(varName string, theVar interface{}) {
    fmt.Print(" ~ " + varName);
    fmt.Printf("(%T)", theVar);
    fmt.Print(" : ", theVar);
    fmt.Println();
}
