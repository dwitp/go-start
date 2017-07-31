package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    ss := strings.Fields(s)
    m := make(map[string]int)
    for i := range ss {
        v := ss[i]
        m[v] = m[v] + 1
    }
    return m
}

func main() {
    wc.Test(WordCount)
}
