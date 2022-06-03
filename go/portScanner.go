// portScanner.go
package main

import (
    "net"
    "os/exec"
    "strconv"
    "strings"
    "time"
)

type PortScanner struct {
    port int
    timeout int
    //method string
    target string
    protocol string
    // more options here
}

type PortScannerResult struct {
    port int
    status bool
}

func ScanPort(data PortScanner) PortScannerResult {
    portString := strconv.Itoa(data.port)
    target := data.target + ":" + portString
    result := PortScannerResult{port: data.port}
    conn, err := net.DialTimeout(data.protocol, target, time.Duration(data.timeout)*time.Millisecond)
    if err != nil {
        if strings.Contains(err.Error(), "too many open files") {
            Printf("[#] ~ Connection overflow non-fatal error -> trying to handle", 3)
            time.Sleep(time.Duration(data.timeout)*time.Millisecond)
            ScanPort(data)
            return result
        } else {
            //fmt.Println(port, "closed")
            result.status = false
            return result
        }
    } else {
        result.status = true
        return result 
    }
    conn.Close()
    return result
}


// func ulimit execs "ulimit -n" shell command to get info about the max number of permited threads 
func ulimit() int64 {
    out, err := exec.Command("ulimit", "-n").Output()    
    if err != nil {
        Printf("[#] ~ Core error while accessing system console\n", 9)
    }    
    s := strings.TrimSpace(string(out))
    i, err := strconv.ParseInt(s, 10, 64)

    if err != nil {
        Printf("[#] ~ Error while trying to parse string\n", 9)
    }    
    return i
}
