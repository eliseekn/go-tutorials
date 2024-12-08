package scraper

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

func Init(query string, resultsLength int) []string {
	c := colly.NewCollector()
	var links []string

	c.OnError(func(_ *colly.Response, err error) {
		panic(err)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		text := e.ChildText("h3")
		link := e.Attr("href")

		if text != "" {
			links = append(links, fmt.Sprintf("[-] %v [%v]", text, formatLink(link)))
		}
	})

	c.Visit("https://google.com/search?hl=en&q=" + formatQuery(query) + "&start=" + fmt.Sprintf("%v", resultsLength))

	return links
}

func formatLink(link string) string {
	link = strings.ReplaceAll(link, "/url?q=", "")

	index := strings.Index(link, "&sa=U&ved=")

	if index != -1 {
		link = link[:index]
	}

	link, err := url.QueryUnescape(link)
	if err != nil {
		panic(err)
	}

	return link
}

func formatQuery(query string) string {
	var result string

	if strings.Contains(query, ":") {
		return url.QueryEscape(query)
	}

	for _, q := range strings.Split(query, " ") {
		result = result + q + "+"
	}

	return result
}
