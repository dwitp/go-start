package main

import (
    "fmt"
)

func dump(args ...interface{}) {
    fmt.Println("-----------------------------------------------------")
    for i := range args {
        fmt.Printf("%+v\n", args[i])
    }
    fmt.Println("-----------------------------------------------------")
    fmt.Println()
}

