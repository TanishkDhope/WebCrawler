package main

import (
	"fmt"
	"time"
)


func main() {
	start:=time.Now()
	myCrawler:=Crawler{
		url: "https://gobyexample.com",
		count: 0,
	}
	myCrawler.crawl(10)
	fmt.Println(myCrawler.visited)
	fmt.Println(time.Since(start))
}
