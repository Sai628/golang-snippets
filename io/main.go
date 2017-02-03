package main

import (
    "os"
    "fmt"
)

func writeToFile(name string, text string) (int, error) {
    file, err := os.Create(name)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    return fmt.Fprintf(file, text)
}

func main() {
    writeToFile("./result.test", "hello golang")
}
