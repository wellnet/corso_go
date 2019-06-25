Waiting Group

/8

package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
    "time"
)

var urls = []string{
    "https://google.com",
    "https://twitter.com",
    "https://google.com",
    "https://twitter.com",
    "https://google.com",
    "https://twitter.com",
    "https://twitter.com",
    "https://google.com",
    "https://twitter.com",
    "https://twitter.com",
    "https://google.com",
    "https://twitter.com",
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return "", err
    }
    wg.Done()
    fmt.Println(resp.Status)
    return resp.Status, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Start!")
    fmt.Println("Urls: ", len(urls))
    var wg sync.WaitGroup
    t1 := time.Now() // time start
    for _, url := range urls {
        wg.Add(1)
        go fetch(url, &wg)
    }

    wg.Wait()
    t2 := time.Now() // end
    fmt.Printf("Time: %v\n", t2.Sub(t1))
    fmt.Printf("Request/second: %v\n", float64(len(urls))/t2.Sub(t1).Seconds())
    fmt.Println("Returning Response")
    fmt.Fprintf(w, "Responses")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":12345", nil))
}

func main() {
    fmt.Println("http://localhost:12345")
    handleRequests()
}

