package main

import (
	"net/http"
	"regexp"
	"sync"
)

func ParseJS(urls []string) {
	var wg sync.WaitGroup
	re := regexp.MustCompile(`/api/[^"' ]+`)

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resp, err := http.Get(u)
			if err != nil {
				return
			}
			defer resp.Body.Close()
			re.FindAll(resp.Body.Read, -1)
		}(url)
	}
	wg.Wait()
}
