package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// sample XML
// <sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
// 	<sitemap>
// 		<loc>https://www.coindesk.com/post-sitemap1.xml</loc>
// 		<lastmod>2014-01-09T16:40:57+00:00</lastmod>
// 	</sitemap>

// SiteMapIndex will contain a slice of Location types
type SiteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

// Location is the url we want to crawl
type Location struct {
	Loc string `xml:"loc"`
}

// Convert the Location struct value to a string
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.coindesk.com/sitemap_index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)
	//	fmt.Println(s.Locations)
	for _, location := range s.Locations {
		fmt.Printf("\n%s", location)
	}
}
