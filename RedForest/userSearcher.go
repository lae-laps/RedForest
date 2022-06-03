package main

import (
    "fmt"
    "time"
    "strconv"
    "net/http"
    "github.com/PuerkitoBio/goquery"
)

// easy sites -> can determine if exists with just status code
// hard sites -> need more advanced specific web scraping for determining if exist

var EASY_SITES = []string{"https://instagram.com/", "https://github.com/"}
var HARD_SITES = []string{"https://twitter.com/"}

type userSearcher struct {
    username string
    //sites []string
}

func (self userSearcher) search() {
    for _, site := range EASY_SITES {
        site = site + self.username
        client := &http.Client{
            Timeout: 20 * time.Second,
        }
        request, err := http.NewRequest("GET", site, nil)
        if err != nil {
            PrintRuntimeError("Error when constructing GET request to : " + site) // construct base request
        }
        request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0") // spoof header to appear as a normal browser
        
        response, err := client.Do(request) // make request

        if err != nil {
            Printf("[400] ", 208)
            fmt.Println(site)
        } else {
            if response.StatusCode >= 200 && response.StatusCode < 300 {
                Printf("[200] ", 10)
                fmt.Println(site)
                //self.sites = append(self.sites, site)
            } else {
                statuscode := strconv.Itoa(response.StatusCode)
                Printf("[" + statuscode + "] ", 208)
                fmt.Println(site)
            }
        }
    }
    for _, site := range HARD_SITES {
        site = site + self.username
        client := &http.Client{
            Timeout: 20 * time.Second,
        }
        request, err := http.NewRequest("GET", site, nil)
        if err != nil {
            PrintRuntimeError("Error when constructing GET request to : " + site) // construct base request
        }
        request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0") // spoof header to appear as a normal browser
        
        response, err := client.Do(request) // make request
        
        if err != nil {
            PrintRuntimeError("Probably an exception - site could not be found at : " + site)
        }


        fmt.Println("uganda")
        fmt.Println(response.Body)

        document, err := goquery.NewDocumentFromReader(response.Body)

        if err != nil {
            PrintRuntimeError("Could not transform request body to a GoQuery object")
        }

        x := document.Find("Hmm...this page doesn’t exist. Try searching for something else.")

        fmt.Println(x)


    }
}

