package main

import (
	"encoding/xml"
	"fmt"

	"github.com/gocolly/colly"
	"github.com/k3a/html2text"
	"github.com/securisec/go-keywords"
)

var sitemapPaths = []string{
	"sitemap.xml",
	"sitemap_index.xml",
	"sitemap-index.xml",
	"sitemap.php",
	"sitemap.txt",
	"sitemap.xml.gz",
	"sitemap",
	"sitemap/sitemap.xml",
	"sitemapindex.xml",
	"sitemap/index.xml",
	"sitemap1.xml",
}

type Sitemap struct {
	Urls []Url `xml:"url"`
}

type Url struct {
	Loc string `xml:"loc"`
}

func fetchUrlsFromDomainname(url string, filter func(url Url) bool) []string {
	c := colly.NewCollector()
	var urls []string

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)

		var sitemap Sitemap

		err := xml.Unmarshal(r.Body, &sitemap)

		if err != nil {
			fmt.Println("Error: %v\n", err)
			return
		}

		for _, url := range sitemap.Urls {
			if filter(url) {
				urls = append(urls, url.Loc)
			}
		}
	})

	for _, sitemapPath := range sitemapPaths {
		c.Visit(fmt.Sprintf("%v/%v", url, sitemapPath))
	}

	return urls
}

func analyseUrl(url string) {
	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		text := html2text.HTML2Text(e.Text)
		keywords, _ := keywords.Extract(string(text))
		fmt.Print(keywords)
	})

	c.Visit(url)
}
