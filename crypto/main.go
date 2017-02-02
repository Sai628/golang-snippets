package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
)

func getMD5Hash(text string) string {
    hasText := md5.New()
    hasText.Write([]byte(text))
    return hex.EncodeToString(hasText.Sum(nil))
}

func main() {
    fmt.Println(getMD5Hash("hello world"))
}
