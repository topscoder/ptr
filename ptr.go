package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    args := os.Args[1:]

    if len(args) != 1 {
        fmt.Println("Usage: ptr <ip_address>")
        return
    }

    ip := args[0]

    ptrDomainName, err := net.LookupAddr(ip)
    if err != nil {
        fmt.Println("No PTR record found for", ip)
        return
    }

    fmt.Println(ptrDomainName)
}
