package main

import (
	"flag"
	"fmt"
	"google-search/scraper"
)

func main() {
	query := flag.String("q", "", "search query")
	length := flag.Int("l", 0, "search results length")

	flag.Parse()

	for i := 0; i <= *length; i += 10 {
		links := scraper.Init(*query, i)

		for _, link := range links {
			fmt.Println(link)
		}
	}
}
