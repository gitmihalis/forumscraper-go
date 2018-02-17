package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// A helper func to pull the href attr from a Token
func getHref(t html.Token) (ok bool, href string) {
	// iterate over all the Token's attributes until we find `href`
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// `bare` return will return the vars ( ok, href )
	return
}

// Extract all http** links from page
func crawl(url string, ch chan string, chFinished chan bool) {
	resp, err := http.Get(url)

	defer func() {
		// Notify when this function finishes
		chFinished <- true
	}()

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
	}

	b := resp.Body
	defer b.Close() // close Body when function returns

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of document, done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if toke is <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, url := getHref(t)
			if !ok {
				continue
			}

			// Check the url begins in http**
			hasProto := strings.Index(url, "http") == 0
			if hasProto {
				ch <- url
			}
		}
	}
}

func main() {
	foundUrls := make(map[string]bool)
	seedUrls := os.Args[1:]

	// Chans
	chUrls := make(chan string)
	chFinished := make(chan bool)

	// Start the crawl process ( concurrently )
	for _, url := range seedUrls {
		go crawl(url, chUrls, chFinished)
	}

	// Subscribe to both channels
	for c := 0; c < len(seedUrls); {
		select {
		case url := <-chUrls:
			foundUrls[url] = true
		case <-chFinished:
			c++
		}
	}

	// Print the rsults....
	fmt.Println("\nFound", len(foundUrls), "unique urls:\n")

	for url, _ := range foundUrls {
		fmt.Println(" - " + url)
	}

	close(chUrls)
}
