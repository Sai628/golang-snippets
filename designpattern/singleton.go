package main

import (
    "sync"
    "fmt"
)

type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}

func main() {
    fmt.Printf("%p\n", GetInstance())
    fmt.Printf("%p\n", GetInstance())
}