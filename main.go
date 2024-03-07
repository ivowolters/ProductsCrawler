package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	urls := fetchProductsFromUrl("https://www.rabobank.nl/")
	for _, url := range urls {
		fmt.Println(url)
		analyseUrl(url)
	}
}
