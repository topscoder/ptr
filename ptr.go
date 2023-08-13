package main

import (
    "fmt"
    "net"
    "os"
    "strings"
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

    // Remove the brackets from the PTR domain name.
    ptrDomainName = strings.Trim(ptrDomainName, "[")
    ptrDomainName = strings.Trim(ptrDomainName, "]")

    // Remove the last character if it is a dot.
    if ptrDomainName[len(ptrDomainName)-1] == '.' {
        ptrDomainName = ptrDomainName[:len(ptrDomainName)-1]
    }

    fmt.Println(ptrDomainName)
}
