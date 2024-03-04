package main

import (
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

func fetchProductsFromUrl(url string) {
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	for _, sitemapPath := range sitemapPaths {
		c.Visit(fmt.Sprintf("%v/%v", url, sitemapPath))
	}
}
