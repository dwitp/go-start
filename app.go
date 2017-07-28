package main

func main() {
    a, b := 1, 2
    dump(a, b)
    var c int
    c = a
    a = b
    b = c
    dump(a, b)
    a, b = b, a
    dump(a, b)
}
