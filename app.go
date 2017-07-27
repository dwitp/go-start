package main

import (
    "./pkg/config"
    "github.com/burntsushi/toml"
)

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
}
