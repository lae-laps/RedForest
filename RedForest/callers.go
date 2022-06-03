// callers.go
package main

import (
    "fmt"
    "regexp"
    "strconv"
//    "strings"
    "time"
)

func temporal(port int) {
}

func PortScannerCaller(sliced []string) {
    
    // TODO: Implement this so it is not total garbage -> use OOP pls

    scanStartTime := time.Now()
    data := PortScanner{}
    if len(sliced) > 1 {
        
        target := sliced[len(sliced) - 1]  
        Printf("[*] ~ Locked target to -> <" + target + ">\n", 6)
       
        match, err := regexp.MatchString(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.){3}(25[0-5]|(2[0-4]|1\d|[1-9]|)\d)$`, target)

        if err != nil {
            Printf("[!] ~ Non-fatal error while trying to apply regexp - defaulting to using raw specified target\n", 3)
        } else {
            if match {
                data.target = target 
            } else {
                Printf("[*] ~ Domain provided -> performing DNS Query...\n", 2)
                //check if result is valid
                ip := DnsQuery(target)
                if len(ip) > 0 {
                    data.target = ip[0]
                } else {
                    Printf("[!] ~ Index out of range -> locking target to domain\n", 4)
                    data.target = target
                }
                Printf("[*] ~ Locked IP target to -> <" + data.target + ">\n", 5)
            }
        } 
    } else {
        Printf("[#] ~ Please specify a target as your last argument\n", 9)
    }

    // iterate over sliced, chech if strings start with - and if true look inside them for specific letters

/*
    var minPort, maxPort int
    for index, content := range sliced {
        if strings.HasPrefix(content, "-") {
            switch content {
                case "-p":
                    if len(sliced) >= index + 2 {
                        var err1, err2 error
                        minPort, err1 = strconv.Atoi(sliced[index + 1])
                        maxPort, err2 = strconv.Atoi(sliced[index + 1])
                        if err1 != nil || err2 != nil {
                            Printf("[!] ~ Non-fatal error in type conversion", 4)
                        }
                    } else {
                        Printf("[-] Specify a port number or range", 9)
                    }
            }
        }
    }
*/
    minPort := 1
    maxPort := 1000
    timeout := 600
    protocol := "tcp"

    data.timeout = timeout
    data.protocol = protocol

    var openPorts []int
    closedPortsCounter := 0


    for i := minPort; i <= maxPort; i++ {
        data.port = i
        
        
        //start := time.Now()
        result := ScanPort(data)
        //total := time.Since(start)

        if result.status == true { // port is open
            //fmt.Printf("Scanning Open port %d took -> %s\n", i, total)
            openPorts = append(openPorts, result.port)
        } else {
            //fmt.Printf("Scanning Closed port %d took -> %s\n", i, total)
            closedPortsCounter ++
        }
    }
    

    Printf("Found " + strconv.Itoa(closedPortsCounter) + " closed ports\n", 4)
    for _, port := range openPorts {
        portStr := strconv.Itoa(port)
        Printf(portStr + ": OPEN\n", 4)
    }

    scanTotalTime := time.Since(scanStartTime)
    scanTotalTimeString := fmt.Sprintf("%s", scanTotalTime.Round(10*time.Millisecond))
    Printf("[*] Scan took " + scanTotalTimeString + "\n", 6)
}

/*
func WebScanCaller(sliced []string) {
    if len(sliced) > 1 {
        target := sliced[len(sliced) - 1]
        getResult := HttpGet(target)
        Printf(getResult.content, 5)
    } else {
        Printf("[#] ~ Please specify a target\n", 9)
    }
}
*/

func userSearcherCaller(sliced []string) {
    if len(sliced) <= 1 {
        PrintUserError("Please specify at least one target to search")
    } else {
        sliced = sliced[1:]
        for _, user := range sliced {
            user := userSearcher{username: user}
            user.search()
        }
    }
}

func DnsLookupCaller(sliced []string) {
    if len(sliced) == 1 {
        PrintUserError("Please specify at least one target")
        
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

