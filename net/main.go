package main

import (
    "errors"
    "fmt"
    "net"
)

func GetLanIP() (string, error) {
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
    ip, err := GetLanIP()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(ip)
    }
}
