package crawler

import (
	"log"
	"net/url"
	"sync"
)

type PageData struct {
	URL   string
	Title string
}

// StartCrawling starts crawling from startURL up to maxDepth levels deep.
// Returns a slice of PageData with results.
func StartCrawling(startURL string, maxDepth int) []PageData {
	visited := make(map[string]bool)
	var mu sync.Mutex

	var results []PageData
	var wg sync.WaitGroup

	var crawl func(string, int)
	crawl = func(currentURL string, depth int) {
		defer wg.Done()
		if depth > maxDepth {
			return
		}

		mu.Lock()
		if visited[currentURL] {
			mu.Unlock()
			return
		}
		visited[currentURL] = true
		mu.Unlock()

		title, err := ParseTitle(currentURL)
		if err != nil {
			log.Printf("Failed to parse %s: %v", currentURL, err)
			return
		}

		mu.Lock()
		results = append(results, PageData{URL: currentURL, Title: title})
		mu.Unlock()

		// Extract URLs to crawl next
		links, err := ExtractLinks(currentURL)
		if err != nil {
			log.Printf("Failed to extract links from %s: %v", currentURL, err)
			return
		}

		baseURL, err := url.Parse(currentURL)
		if err != nil {
			return
		}

		for _, link := range links {
			absURL := toAbsoluteURL(link, baseURL)
			if absURL == "" {
				continue
			}

			wg.Add(1)
			go crawl(absURL, depth+1)
		}
	}

	wg.Add(1)
	go crawl(startURL, 1)

	wg.Wait()
	return results
}

func toAbsoluteURL(href string, base *url.URL) string {
	parsed, err := url.Parse(href)
	if err != nil {
		return ""
	}
	abs := base.ResolveReference(parsed)
	return abs.String()
}
