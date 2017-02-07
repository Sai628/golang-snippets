package main

import (
    "os"
    "fmt"
    "path/filepath"
    "io/ioutil"
)

func writeToFile(name string, text string) (int, error) {
    file, err := os.Create(name)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    return fmt.Fprintf(file, text)
}

func renameFile(dir, oldName, newName string) error {
    oldPath := filepath.Join(dir, oldName)
    newPath := filepath.Join(dir, newName)
    return os.Rename(oldPath, newPath)
}

func listDir(dir string) {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    for _, f := range files {
        fmt.Println(f.Name())
    }
}

func main() {
    writeToFile("./result.test", "hello golang")
    renameFile("./", "result.test", "result2.test")
    listDir("./io/")
}
