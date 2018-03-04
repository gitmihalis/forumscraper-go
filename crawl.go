package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// sample XML
// <sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
// 	<sitemap>
// 		<loc>https://www.coindesk.com/post-sitemap1.xml</loc>
// 		<lastmod>2014-01-09T16:40:57+00:00</lastmod>
// 	</sitemap>

// SiteMapIndex will contain a slice of Location types
type SiteMapIndex struct {
	Locations []string `xml:"url>loc"`
}

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

func main() {
	var s SiteMapIndex

	resp, _ := http.Get("https://bitcoinforum.com/sitemap/?xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	// Visit the locations on the sitemap
	for _, location := range s.Locations {
		resp, _ := http.Get(location)
		body := resp.Body
		defer body.Close()

		z := html.NewTokenizer(body)

		for {
			// repeatedly call z.Next() which parses the next token and returns it's type
			tt := z.Next()

			switch {
			case tt == html.ErrorToken:
				return
			case tt == html.StartTagToken:
				// Process the current token.
				t := z.Token()

				// Check is token <a> tag?
				isAnchor := t.Data == "a"
				if !isAnchor {
					continue
				}

				// Extract href
				ok, url := getHref(t)
				if !ok {
					continue
				}

				// Check the url begins in http**
				hasProto := strings.Index(url, "http") == 0
				if hasProto {
					fmt.Println(url)
				}
			}

		}

	}

	// TODO: Build a Map!
}
