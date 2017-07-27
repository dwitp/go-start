package main

import (
    "./pkg/config"
    "github.com/burntsushi/toml"
    "time"
)

func say(s string) {
    for i := 0; i < 3; i++ {
        time.Sleep(1000 * time.Millisecond)
        printVar("s", s)
    }
}

func main() {
    confData, err := FindAndReadConfig("config.toml")
    if err != nil {
        printVar("err", err)
    }

    var config config.Config
    if _, err := toml.Decode(confData, &config); err != nil {
        printVar("err", err)
    }

    printVar("config", config)

    go say("world")
    say("hello")
}
