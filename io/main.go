package main

import (
    "bytes"
    "fmt"
    "net/http"
)

func GetResponse(url string) (string, error) {
    response, err := http.Get(url)
    if err != nil {
        return "", err
    }
    buf := new(bytes.Buffer)
    buf.ReadFrom(response.Body)
    return buf.String(), nil
}

func main() {
    content, _ := GetResponse("http://httpbin.org/user-agent")
    fmt.Println(content)

    // {
    //   "user-agent": "Go-http-client/1.1"
    // }
}
