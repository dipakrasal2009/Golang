// dns is a simple command line tool to lookup the ip address of a host;
// it prints the first ipv4 and ipv6 addresses it finds, or "none" if none are found.
package main

import (
    "fmt"
    "log"
    "net"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        log.Infof("%s: usage: <host>", os.Args[0])
        log.Fatalf("expected exactly one argument; got %d", len(os.Args)-1)
    }
    host := os.Args[1]
    ips, err := net.LookupIP(host)
    if err != nil {
        log.Fatalf("lookup ip: %s: %v", host, err)
    }
    if len(ips) == 0 {
        log.Fatalf("no ips found for %s", host) // this should never happen, but just in case
    }
    // print the first ipv4 we find
    for _, ip := range ips {
        if ip.To4() != nil {
            fmt.Println(ip)
            goto IPV6 // goto considered awesome
        }
    }
    fmt.Printf("none\n") // only print "none" if we don't find any ipv4 addresses

IPV6: // print the first ipv6 we find
    for _, ip := range ips {
        if ip.To4() == nil {
            fmt.Println(ip) // we don't need to check for nil here, since we know we have at least one ip address
            return
        }
    }
    fmt.Printf("none\n")
}


