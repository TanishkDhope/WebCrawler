package main

import (
	"sync"
	//"fmt"
)


type Crawler struct{
	url string
	count int
	limit int
	visited Stack[string]
	mu sync.Mutex
}

/*
func (c *Crawler) incr(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
*/
func (c *Crawler) crawl(limit int){


	c.visited=Stack[string]{
			data: make([]string,0),
			len: 0,
		}
	
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
                // Stop condition
                c.mu.Lock()
                if c.count >= c.limit {
                    c.mu.Unlock()
                    return
                }
                c.mu.Unlock()

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
				c.visited.push(url)

                c.mu.Unlock()

                // Visit the URL
                c.visit(url, &links)

            }
        }()
    }

    wg.Wait()
}
