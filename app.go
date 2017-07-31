package main

import (
    "github.com/dwitp/golang/src/u"
)

func main() {
    m := make(map[string]int)

    m["Answer"] = 42
    u.D("The value:", m["Answer"])

    m["Answer"] = 48
    u.D("The value:", m["Answer"])

    delete(m, "Answer")
    u.D("The value:", m["Answer"])

    v, ok := m["Answer"]
    u.D("The value:", v, "Present?", ok)
}
