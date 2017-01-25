package main

import (
	"sync"
	"log"
)

func main() {
    wg := &sync.WaitGroup{}

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(wg *sync.WaitGroup, i int) {
            log.Printf("i:%d", i)
            wg.Done()
        }(wg, i)
    }

    wg.Wait()
    log.Println("exit")
}
