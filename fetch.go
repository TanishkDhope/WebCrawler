package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"net/http"
)

func (c *Crawler) visit(url string, links *Stack[string]) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			panic(err)
		}
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			if href, exists := s.Attr("href"); exists && strings.HasPrefix(href, "https") {
				links.mu.Lock()
				links.push(href)
				links.mu.Unlock()
			}	
		})

	}
}