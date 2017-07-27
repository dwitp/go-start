package main

import "fmt"

func printVar(varName string, theVar interface{}) {
    fmt.Print(" ~ " + varName);
    fmt.Printf("(%T)", theVar);
    fmt.Print(" : ", theVar);
    fmt.Println();
}

