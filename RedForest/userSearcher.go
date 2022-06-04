package main

import (
    "fmt"
    "time"
    "strings"
    "strconv"
    "net/http"
)

// easy sites -> can determine if exists with just status code
// hard sites -> need more advanced specific web scraping for determining if exist - custom function for each one

var SITES = []string{"https://www.reddit.com/user/$", "https://www.pinterest.com/$", "https://open.spotify.com/user/$", "https://youtube.com/$", "https://twitch.tv/$", "https://twitter.com/$", "https://instagram.com/$", "https://t.me/$", "https://github.com/$", "https://www.flickr.com/photos/$", "https://$.blogspot.com", "https://plus.google.com/+$/posts", "https://$.wordpress.com", "https://steamcommunity.com/id/$", "https://vimeo.com/$", "https://soundcloud.com/$", "https://disqus.com/$", "https://medium.com/@$", "https://$.devianart.com", "https://about.me/$", "https://www.mixcloud.com/$", "https://bitbucket.com/$", "https://cash.me/$", "https://behance.net/$", "https://pastebin.com/u/$", "https://$.newgrounds.com", "https://wattpad.com/user/$", "https://creativemarket.com/$", "https://$.contently.com", "https://buzzfeed.com/$", "https://tripadvisor.com/members/$", "https://codementor.io/$", "https://ebay.com/usr/$", "https://$.slack.com", "https://www.devianart/com/$", "https://gitlab.com/$", "https://hackthebox.com/profile/$"}

type userSearcher struct {
    username string
    //sites []string
}

func (self userSearcher) search() {
    counter := 0
    startTime := time.Now()
    for _, site := range SITES {
        counter ++
        site = strings.Replace(site, "$", self.username, 1)
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
    totalTime := time.Since(startTime)
    PrintBold("[!] ", 84)
    fmt.Printf("Scan took a total of %s seconds\n", totalTime.Round(10*time.Millisecond))
    PrintBold("[!] ", 21)
    fmt.Printf("Scanned a total of %d sites\n", counter)
}

