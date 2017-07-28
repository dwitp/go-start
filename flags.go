package main

import (
    "flag"
    "fmt"
    "github.com/dwitp/go-start/src/util"
    "os"
)

const (
    DefaultConfigPath = ""
)

var ConfigPaths []string = []string{
    ".",
    os.Getenv("HOME"),
    "/etc",
}
var configFilePath string

func init() {
    flag.StringVar(&configFilePath,
        "config", DefaultConfigPath, "path to config file",
    )
}

func FindAndReadConfig(path string) (string, error) {
    var content []byte
    var err error

    if path != "" {
        content, err = util.ReadConfigFile(path)
    } else {
        content, err = readDefaultConfigPath("config.toml")
    }

    return string(content), err
}

func readDefaultConfigPath(fileName string) ([]byte, error) {
    var content []byte
    var err error

    for _, configPath := range ConfigPaths {
        content, err = util.ReadConfigFile(configPath + "/" + fileName)
        if err == nil {
            fmt.Println("Find log files on path:", configPath)
            break
        }
    }

    return content, err
}
