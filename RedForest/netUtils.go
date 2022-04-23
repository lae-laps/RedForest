// netUtils.go
package main

import (
	"net"
)

type DnsQueryResult []string

func DnsQuery(target string) DnsQueryResult  {
    var result DnsQueryResult
    ip, err := net.LookupIP(target)
    if err != nil {
        Printf("Failed to make DNS request\n", 6)
    } else {
        for _, ip := range ip {
            result = append(result, ip.String())
        }
    }
    return result
}



