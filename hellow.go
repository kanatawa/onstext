// crawler.go
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	urls := []string{"http://example.com", "http://example.org", "http://example.net"}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(url, resp.Status)
		}(url)
	}
	wg.Wait()
}
