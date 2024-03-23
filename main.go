package main

import (
	"fmt"
	"productcrawler/db"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	urls := fetchUrlsFromDomainname(
		"https://www.exact.com/",
		func(url Url) bool { return strings.Contains(url.Loc, "product") },
	)
	for _, url := range urls {
		fmt.Println(url)
		result := analyseUrl(url)

		if result.Title == "" {
			fmt.Print("Empty body")
			continue
		}

		db.SaveProduct(db.ProductDto{
			Url:      url,
			Title:    result.Title,
			Keywords: result.Keywords,
		})

		fmt.Println(result)
	}
}
