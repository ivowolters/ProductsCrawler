package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	urls := fetchProductsFromUrl(
		"https://www.exact.com/",
		func(url Url) bool { return strings.Contains(url.Loc, "product") },
	)
	for _, url := range urls {
		fmt.Println(url)
		analyseUrl(url)
	}
}
