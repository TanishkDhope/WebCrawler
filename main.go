package main

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

func main() {
	db,err:=bolt.Open("index.db",0600, nil)
	if err!=nil{
		panic(err)
	}

	defer db.Close()

	myCrawler:=Crawler{
		url: "https://gobyexample.com",
		count: 0,
	}
	for{
		fmt.Println("\n1.Crawl\n2.Query\n3.Exit")
		var choice int
		fmt.Printf("Enter Your Choice: ")
		fmt.Scan(&choice)
		switch choice{
		case 1:
			var url string 
			var limit int
			fmt.Println("Enter url to Crawl and limit: ")
			fmt.Scan(&url,&limit)
			myCrawler.url=url
			start:=time.Now()
			myCrawler.crawl(limit,db)
			fmt.Println(time.Since(start))
		case 2:
		var query string
		fmt.Printf("Enter Search words: ")
		fmt.Scan(&query)
		results:=search(db,query,10)
		fmt.Println(results)
		default:
			return
		}
	}
}
