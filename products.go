package main

import (
	"encoding/xml"
	"fmt"

	"github.com/gocolly/colly"
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

func fetchProductsFromUrl(url string) {
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)

		var sitemap Sitemap

		err := xml.Unmarshal(r.Body, &sitemap)

		if err != nil {
			fmt.Println("Error: %v\n", err)
			return
		}
	})

	for _, sitemapPath := range sitemapPaths {
		c.Visit(fmt.Sprintf("%v/%v", url, sitemapPath))
	}
}
