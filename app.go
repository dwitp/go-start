package main

import (
)

func main() {
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    dump("names", names)

    a := names[0:2]
    b := names[1:3]
    dump(a, b)

    b[0] = "XXX"
    dump(a, b)
    dump(names)
}
