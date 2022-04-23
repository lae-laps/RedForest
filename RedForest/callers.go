// callers.go
package main

import (
    "fmt"
    "strings"
)

func DnsLookupCaller(expression string) {
    sliced := strings.Fields(expression)
    if len(sliced) == 1 {
        Printf("[#] ~ Please specify at least one target\n", 9)
        
    } else {
        Printf("\nResults for DNS query:\n======================\n", 8) 
        for i, content := range sliced {
            if i > 0 {
                Printf("\nTarget: " + content + "\n", 10)
                result := DnsQuery(content)
                for _, ip := range result {
                    Printf("|> " + ip + "\n", 4)
                }
            }
        }
        fmt.Print("\n")
    }
}

