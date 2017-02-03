package main

import (
    "errors"
    "fmt"
    "net"
    "net/http"
    "bytes"
)

func getHTTPResponse(url string) (string, error) {
    response, err := http.Get(url)
    if err != nil {
        return "", err
    }
    buf := new(bytes.Buffer)
    buf.ReadFrom(response.Body)
    return buf.String(), nil
}

func getLanIP() (string, error) {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return "", err
    }

    for _, a := range addrs {
        if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String(), nil
            }
        }
    }

    return "", errors.New("not found lan ip")
}

func main() {

    // test getHTTPResponse
    content, _ := getHTTPResponse("http://httpbin.org/user-agent")
    fmt.Println(content)
    // {
    //   "user-agent": "Go-http-client/1.1"
    // }

    // test getLanIP
    ip, err := getLanIP()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(ip)
    }
}
