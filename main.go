package main

import (
	"context"
	"fmt"
	"time"
	"net/http"
) 

type Result struct {
	Status string
	Urlstring string
	Error error
}

func main() {
    fmt.Println("Hello, world")

	//created context with timeout
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	//created a buffered channel
	ch := make(chan Result,5)

	urls := []string{
        "https://google.com",
        "https://github.com",
        "https://golang.org",
		"https://iitkgp.com",
		"https://facebook.com",
		"https://apple.com",
    }

	for _,url := range urls {
		go func(url string){
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil{
				ch <- Result{"failed",url,err}
				return
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil{
				ch <- Result{"failed", url, err}
				return
			}
			ch <- Result{"ok",url,err}
			resp.Body.Close()
		}(url)
	}

	for range len(urls){
		result := <- ch
		fmt.Printf("%s %s %s \n", result.Status, result.Urlstring, result.Error)
	}
}