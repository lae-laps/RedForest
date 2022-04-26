// callers.go
package main

import (
    "fmt"
    "strconv"
   // "time"
)

func temporal(port int) {
}

func PortScannerCaller(sliced []string) {
    data := PortScanner{}
    if len(sliced) > 1 {
        data.target = sliced[len(sliced) - 1]  
        Printf("[*] ~ Locked target to -> <" + data.target + ">\n", 6)
    } else {
        Printf("[#] ~ Please specify a target as your last argument\n", 9)
    }

    // iterate over sliced, chech if strings start with - and if true look inside them for specific letters

    minPort := 1
    maxPort := 65535 
    timeout := 600
    protocol := "tcp"

    data.timeout = timeout
    data.protocol = protocol

    //var openPorts []int
    //closedPortsCounter := 0

    for i := minPort; i <= maxPort; i++ {
//        Printf("-", 5)
        data.port = i
        
        
      //  start := time.Now()
        //result := ScanPort(data)
        str := strconv.Itoa(i)
        Printf(str + "\n", 6)
    //    total := time.Since(start)

    /*
        if result.status == true { // port is open
            fmt.Printf("Scanning Open port %d took -> %s\n", i, total)
            openPorts = append(openPorts, result.port)
        } else {
            //fmt.Printf("Scanning Closed port %d took -> %s\n", i, total)
            closedPortsCounter ++
        }
    }

    closedPortsCounterStr := strconv.Itoa(closedPortsCounter)
    Printf("Found " + closedPortsCounterStr + " closed ports\n", 4)
    for _, port := range openPorts {
        portStr := strconv.Itoa(port)
        Printf(portStr + ": OPEN\n", 4)
    }
*/
    }
}

func DnsLookupCaller(sliced []string) {
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

