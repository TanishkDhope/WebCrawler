package main

import (
	"slices"
	"fmt"
	"net/http"

	//"bufio"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var Links []string

func crawl(url string) {
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
			if href, exists := s.Attr("href"); exists && strings.HasPrefix(href, "http") {
				if !slices.Contains(Links,href) {
					Links = append(Links, href)
				}
			}
		})
		fmt.Println(Links)
	}
}


func main() {
	crawl("https://gobyexample.com")

}
