package main

import (
	"sync"
	"github.com/boltdb/bolt"
)



type Crawler struct{
	url string
	count int
	limit int
	visited map[string]bool
	mu sync.Mutex
}


func (c *Crawler) crawl(limit int,db *bolt.DB){

	c.visited=make(map[string]bool)

	c.limit=limit
	var wg sync.WaitGroup

	links:=Stack[string]{
		data: make([]string, 0),
		len: 0,
	}
    
	links.push(c.url)
	  workerCount := 20

    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()

            for {
                // Try to pop a URL
                links.mu.Lock()
                if links.len == 0 {
                    links.mu.Unlock()
                    continue // wait for other workers to push more links
                }
                url := links.pop()
                links.mu.Unlock()

                // Safe increment before visiting
                c.mu.Lock()
                if c.count >= c.limit {
                    c.mu.Unlock()
                    return
                }
                c.count++
				c.visited[url]=true

                c.mu.Unlock()

                // Visit the URL
                c.visit(db,url, &links)

            }
        }()
    }

    wg.Wait()
}
