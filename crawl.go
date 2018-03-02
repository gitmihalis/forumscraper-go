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
	Locations []string `xml:"sitemap>loc"`
}

// Pages will contain locations of forum pages
type Pages struct {
	Pages []string `xml:"url>loc"`
}

func main() {
	var s SiteMapIndex
	var p Pages
	var reqCount int
	
	resp, _ := http.Get("https://bitcointalk.org/sitemap.php")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	
	// Visit the locations on the sitemap
	for _, location := range s.Locations {
		fmt.Println("Crawling...", reqCount)
		reqCount++
		resp, _ := http.Get(location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &p)
		resp.Body.Close()
	}

	fmt.Println(p, reqCount)

	// TODO: Build a Map!
}
