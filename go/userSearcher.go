package main

import (
    "fmt"
    "net/http"
)

var SITES []string = ["https://twitter.com"]string

type userSearcher struct {
    usernames []string
    sites []string
}

func (self userSearcher) parse(sliced []string) {
    if len(sliced) <= 1 {
        Printf("[#] ~ Please specify at least one target\n", 4)
    } else {
        for _, user := range self.usernames {
            self.usernames = append(self.usernames, user)
        }
    }
}

func (self userSearcher) search() {
    for _, username := range self.usernames {
        for _, site := range SITES {
            _, err := http.Get(site + username)
            if err == nil {
                self.sites = append(self.sites, site)
            }
        }
    }
}

func (self userSearcher) printResults() {
    fmt.Println(self.sites)
}

