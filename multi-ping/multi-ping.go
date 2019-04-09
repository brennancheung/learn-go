package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	domains := []string{"google.com", "facebook.com", "instagram.com", "amazon.com", "yahoo.com"}

	// startTime := time.Now()
	ch := make(chan string)

	for _, domain := range domains {
		fmt.Println(domain)
		go fetch(domain, ch)
	}

	for range domains {
		fmt.Println(<-ch)
	}
}

func fetch(domain string, ch chan<- string) {
	startTime := time.Now()
	url := "http://" + domain
	_, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	secs := time.Since(startTime).Seconds()
	ch <- fmt.Sprintf("%.2f %s", secs, url)
}
