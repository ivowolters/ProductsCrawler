package main

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
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

type AnalyseUrlResponse struct {
	Title    string
	Keywords []string
}

func fetchUrlsFromDomainname(domainName string, filter func(url Url) bool) []string {
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
		c.Visit(fmt.Sprintf("%v/%v", domainName, sitemapPath))
	}

	return urls
}

func analyseUrl(url string) AnalyseUrlResponse {
	c := colly.NewCollector()
	var textSb strings.Builder
	var titleSb strings.Builder

	c.OnHTML("body", func(e *colly.HTMLElement) {
		contentTags := []string{"h1", "h2", "h3", "p"}

		for _, tag := range contentTags {
			e.ForEach(tag, func(i int, h *colly.HTMLElement) {
				textSb.WriteString(fmt.Sprintf("%v ", h.Text))
			})
		}

		e.ForEach("h1", func(i int, h *colly.HTMLElement) {
			titleSb.WriteString(fmt.Sprintf("%v ", h.Text))
		})
	})

	c.Visit(url)

	text := strings.TrimSpace(textSb.String())
	keywords, _ := keywords.Extract(string(text))

	return AnalyseUrlResponse{
		Title:    strings.TrimSpace(titleSb.String()),
		Keywords: keywords,
	}
}
