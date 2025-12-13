package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"net/http"
	"net/url"
)

func normalizeUrl(raw_url string, base *url.URL) string {
	u,err:=url.Parse(raw_url)
	if err!=nil{
		panic(err)
	}

	if base!=nil{
		u=base.ResolveReference(u)
	}

	u.Scheme=strings.ToLower(u.Scheme)
	u.Host=strings.ToLower(u.Host)
	
	u.Fragment=""

	if u.Path==""{
		u.Path="/"
	}

	// Normalize trailing slash (choose one rule) 
	if u.Path != "/" { 
		u.Path = strings.TrimRight(u.Path, "/") 
	}

	return u.String()
}

func (c *Crawler) visit(base string, links *Stack[string]) {
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
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			if href, exists := s.Attr("href"); exists {
				var site string
				u,_:=url.Parse(base)
				site=normalizeUrl(href, u)
			
				links.mu.Lock()
				links.push(site)
				links.mu.Unlock()
				
			}	
		})

	}
}