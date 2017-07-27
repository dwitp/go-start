package main

import "fmt"

func printVar(varName string, theVar interface{}) {
    fmt.Println("-----------------------------------------------------")
    fmt.Print(varName)
    fmt.Printf(" <%T>", theVar)
    fmt.Println()
    fmt.Printf("%+v\n", theVar)
    fmt.Println("-----------------------------------------------------")
    fmt.Println();
}

