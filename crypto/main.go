package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "crypto/sha1"
    "crypto/sha256"
)

func getMD5Hash(text string) string {
    hashText := md5.New()
    hashText.Write([]byte(text))
    return hex.EncodeToString(hashText.Sum(nil))
}

func getSHA1Hash(text string) string {
    hashText := sha1.New()
    hashText.Write([]byte(text))
    return hex.EncodeToString(hashText.Sum(nil))
}

func getSHA256Hash(text string) string {
    hashText := sha256.New()
    hashText.Write([]byte(text))
    return hex.EncodeToString(hashText.Sum(nil))
}

func main() {
    fmt.Println(getMD5Hash("hello world"))  // 5eb63bbbe01eeed093cb22bb8f5acdc3
    fmt.Println(getSHA1Hash("hello world"))  // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
    fmt.Println(getSHA256Hash("hello world"))  // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
}
