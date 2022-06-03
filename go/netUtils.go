// netUtils.go
package main

import (
	"net"
//    "net/http"
  //  "io"
)

type DnsQueryResult []string

func DnsQuery(target string) DnsQueryResult {
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

/*
type HttpResult struct {
    content string
}

func HttpGet(target string) HttpResult {
    var result HttpResult

    response, err := http.Get(target)
    if err != nil {
        Printf("[#] ~ Error while performing http GET request\n", 9)
    }

    body, err := io.ReadAll(response.Body)
    if err != nil {
        Printf("[#] ~ Error while reading response body to string\n", 9)
    }
  
    bodyString := string(response.Body)

    result.content = bodyString
    return result
} 
*/
