package main

import (
    "sync"
    "fmt"
)

func main() {
    wg := &sync.WaitGroup{}

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(wg *sync.WaitGroup, i int) {
            fmt.Printf("i:%d", i)
            wg.Done()
        }(wg, i)
    }

    wg.Wait()
    fmt.Println("exit")
}
