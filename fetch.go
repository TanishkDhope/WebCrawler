package main

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/boltdb/bolt"
)

func normalizeUrl(raw_url string, base *url.URL) (string, error) {
	u, err := url.Parse(raw_url)
	if err != nil {
		return "", err
	}

	if base != nil {
		u = base.ResolveReference(u)
	}

	u.Scheme = strings.ToLower(u.Scheme)
	u.Host = strings.ToLower(u.Host)

	u.Fragment = ""

	if u.Path == "" {
		u.Path = "/"
	}

	// Normalize trailing slash (choose one rule)
	if u.Path != "/" {
		u.Path = strings.TrimRight(u.Path, "/")
	}

	return u.String(), nil
}

func (c *Crawler) visit(db *bolt.DB, base string, links *Stack[string]) {
	res, err := http.Get(base)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			panic(err)
		}
		doc.Find("nav, footer, aside, script, style").Remove()

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			if href, exists := s.Attr("href"); exists {
				var site string
				baseURL, _ := url.Parse(base)
				site, err := normalizeUrl(href, baseURL)
				if err != nil {
					return
				}

				links.mu.Lock()
				links.push(site)
				links.mu.Unlock()

			}
		})

		title := doc.Find("title").Text()
		body := doc.Find("body").Text()
		tokens := tokenize(title + " " + body)
		totalTokens := len(tokens)

		addDocument(db,base,totalTokens)

		for _, token := range tokens {
			addToIndex(db, token, base)
		}

	}
}
