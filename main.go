package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"strconv"
	"sync"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/temoto/robotstxt"
)

type DocPage struct {
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
}

func createDirectories(baseURL string) (string, error) {
	// Create base directory
	timestamp := strings.ReplaceAll(strings.ReplaceAll(baseURL, "https://", ""), "/", "-")
	dirPath := filepath.Join("generated_docs", timestamp)
	
	if err := os.MkdirAll(filepath.Join(dirPath, "markdown"), 0755); err != nil {
		return "", err
	}
	
	return dirPath, nil
}

func normalizeURL(url string) string {
	return strings.TrimSuffix(url, "/")
}

func extractContent(s *goquery.Selection) (string, error) {
	// Create a new converter
	converter := md.NewConverter("", true, nil)

	// Get the HTML content
	html, err := s.Html()
	if err != nil {
		return "", err
	}

	// Convert HTML to Markdown
	markdown, err := converter.ConvertString(html)
	if err != nil {
		return "", err
	}

	// Clean up the markdown
	markdown = strings.ReplaceAll(markdown, "\n\n\n", "\n\n") // Remove excessive newlines
	
	// Ensure content ends with a newline
	if !strings.HasSuffix(markdown, "\n") {
		markdown += "\n"
	}
	
	return markdown, nil
}

func sanitizeFilename(title string) string {
	// Convert to lowercase
	name := strings.ToLower(title)
	
	// Replace special characters with underscores
	name = strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return r
		case r >= '0' && r <= '9':
			return r
		case r == ' ' || r == '-':
			return '-'
		default:
			return '_'
		}
	}, name)
	
	// Replace multiple dashes/underscores with single ones
	name = strings.ReplaceAll(name, "--", "-")
	name = strings.ReplaceAll(name, "__", "_")
	
	return name + ".md"
}

func formatMarkdown(title, content string) string {
	var formatted strings.Builder

	// Add metadata section
	formatted.WriteString("---\n")
	formatted.WriteString(fmt.Sprintf("title: %s\n", title))
	formatted.WriteString(fmt.Sprintf("source: Remotion Documentation\n"))
	formatted.WriteString(fmt.Sprintf("last_updated: %s\n", time.Now().Format("2006-01-02")))
	formatted.WriteString("---\n\n")

	// Add title if not present
	if !strings.HasPrefix(content, "# ") {
		formatted.WriteString(fmt.Sprintf("# %s\n\n", title))
	}

	// Add the main content
	formatted.WriteString(content)

	return formatted.String()
}

type RateLimiter struct {
	ticker *time.Ticker
}

func NewRateLimiter(delay time.Duration) *RateLimiter {
	return &RateLimiter{
		ticker: time.NewTicker(delay),
	}
}

func (r *RateLimiter) Wait() {
	<-r.ticker.C
}

func (r *RateLimiter) Stop() {
	r.ticker.Stop()
}

func crawlURL(baseURL string) ([]DocPage, error) {
	// Check robots.txt first
	fmt.Println("Checking robots.txt...")
	robotsURL := "https://www.remotion.dev/robots.txt"
	resp, err := http.Get(robotsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch robots.txt: %v", err)
	}
	defer resp.Body.Close()

	robots, err := robotstxt.FromResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse robots.txt: %v", err)
	}

	group := robots.FindGroup("*")
	if group != nil && !group.Test("/docs") {
		return nil, fmt.Errorf("crawling /docs is not allowed by robots.txt")
	}

	// Get crawl delay from robots.txt, default to 1 second if not specified
	crawlDelay := 1 * time.Second
	if group != nil && group.CrawlDelay > 0 {
		crawlDelay = time.Duration(group.CrawlDelay) * time.Second
	}

	// Create rate limiter
	rateLimiter := NewRateLimiter(crawlDelay)
	defer rateLimiter.Stop()

	// Normalize the input URL
	baseURL = normalizeURL(baseURL)
	fmt.Printf("\nStarting crawl from: %s\n", baseURL)
	fmt.Printf("Rate limit: %v between requests\n", crawlDelay)

	// Default to 4 workers as it provides optimal parallelization while staying
	// well within typical browser connection limits. Combined with rate limiting,
	// this gives good performance without overwhelming the target server.
	workers := 4
	if len(os.Args) > 2 {
		if w, err := strconv.Atoi(os.Args[2]); err == nil && w > 0 {
			workers = w
		}
	}

	fmt.Printf("Number of workers: %d\n\n", workers)

	var pages []DocPage
	processedLinks := make(map[string]bool)
	queuedLinks := make(map[string]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Create a channel for new URLs to crawl
	urlQueue := make(chan string, 1000)
	done := make(chan struct{})
	
	// Add initial URL
	urlQueue <- baseURL
	queuedLinks[baseURL] = true

	// Stats
	processedCount := 0
	queuedCount := 1
	lastProcessedCount := 0
	lastQueuedCount := 1

	// Progress reporting goroutine
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				mu.Lock()
				currentProcessed := processedCount
				currentQueued := queuedCount
				newProcessed := currentProcessed - lastProcessedCount
				newQueued := currentQueued - lastQueuedCount
				fmt.Printf("\nProgress Update:\n")
				fmt.Printf("- Pages processed: %d (+ %d new)\n", currentProcessed, newProcessed)
				fmt.Printf("- Pages queued: %d (+ %d new)\n", currentQueued, newQueued)
				fmt.Printf("- Pages found: %d\n", len(pages))
				if newProcessed == 0 && newQueued == 0 {
					fmt.Printf("(Crawler is still working, but no new pages in the last 5 seconds)\n")
				}
				lastProcessedCount = currentProcessed
				lastQueuedCount = currentQueued
				mu.Unlock()
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	// Create worker goroutines
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case url, ok := <-urlQueue:
					if !ok {
						return
					}

					mu.Lock()
					if processedLinks[url] {
						mu.Unlock()
						continue
					}
					processedLinks[url] = true
					processedCount++
					mu.Unlock()

					// Wait for rate limiter before making request
					rateLimiter.Wait()

					newPages, newLinks, err := crawlPage(url)
					if err != nil {
						fmt.Printf("Error crawling %s: %v\n", url, err)
						continue
					}

					mu.Lock()
					pages = append(pages, newPages...)
					
					// Add new links to the queue
					for _, link := range newLinks {
						if !processedLinks[link] && !queuedLinks[link] {
							select {
							case urlQueue <- link:
								queuedLinks[link] = true
								queuedCount++
							default:
								fmt.Printf("Queue full, skipping URL: %s\n", link)
							}
						}
					}
					mu.Unlock()

				case <-done:
					return
				}
			}
		}(i)
	}

	// Monitor for completion (no progress for 15 seconds)
	go func() {
		noProgressCount := 0
		lastCount := 0
		checkTicker := time.NewTicker(5 * time.Second)
		defer checkTicker.Stop()

		for {
			select {
			case <-checkTicker.C:
				mu.Lock()
				currentCount := len(pages)
				if currentCount == lastCount {
					noProgressCount++
					if noProgressCount >= 3 { // No progress for 15 seconds
						fmt.Println("\nNo new pages found for 15 seconds, finishing crawl...")
						close(urlQueue)
						close(done)
						mu.Unlock()
						return
					}
				} else {
					noProgressCount = 0
				}
				lastCount = currentCount
				mu.Unlock()
			}
		}
	}()

	// Wait for all workers to finish
	wg.Wait()
	fmt.Printf("\nCrawl completed!\n")
	fmt.Printf("Total pages processed: %d\n", processedCount)
	fmt.Printf("Total pages found: %d\n", len(pages))

	return pages, nil
}

func crawlPage(url string) ([]DocPage, []string, error) {
	var pages []DocPage
	var links []string

	// Get the webpage
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, fmt.Errorf("HTTP GET error: %v", err)
	}
	defer resp.Body.Close()

	// Check if it's a 404 page
	if resp.StatusCode == 404 {
		return nil, nil, fmt.Errorf("404 page not found")
	}

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("HTML parsing error: %v", err)
	}

	// Process the page
	title := doc.Find("h1").First().Text()
	if title == "" {
		title = doc.Find("title").First().Text()
	}

	// Find the main content
	var markdown string
	mainContent := doc.Find("article, .markdown, .content, main, #main-content").First()
	if mainContent.Length() > 0 {
		content, err := extractContent(mainContent)
		if err != nil {
			return nil, nil, fmt.Errorf("content extraction error: %v", err)
		}
		markdown = formatMarkdown(title, content)
	}

	if title != "" && markdown != "" && !strings.Contains(title, "Page Not Found") {
		pages = append(pages, DocPage{
			Title:    title,
			Markdown: markdown,
		})
	}

	baseURL := "https://www.remotion.dev"

	// Find all documentation links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		// Skip anchors and non-doc pages
		if strings.HasPrefix(href, "#") {
			return
		}

		// Handle absolute URLs
		if strings.HasPrefix(href, "http") {
			if strings.Contains(href, "remotion.dev/docs") {
				links = append(links, href)
			}
			return
		}

		// Handle relative URLs
		if strings.HasPrefix(href, "/") {
			href = strings.TrimPrefix(href, "/")
		}

		// Skip non-documentation pages
		if !strings.Contains(href, "docs") {
			return
		}

		// Construct absolute URL
		absoluteURL := fmt.Sprintf("%s/%s", baseURL, href)
		links = append(links, absoluteURL)
	})

	return pages, links, nil
}

func writeResults(pages []DocPage, dirPath string) error {
	// Write JSON file
	jsonData, err := json.MarshalIndent(pages, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(dirPath, "crawl_results.json"), jsonData, 0644); err != nil {
		return err
	}

	// Write individual markdown files
	for _, page := range pages {
		filename := sanitizeFilename(page.Title)
		filepath := filepath.Join(dirPath, "markdown", filename)
		if err := os.WriteFile(filepath, []byte(page.Markdown), 0644); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var url string
	if len(os.Args) > 1 {
		url = os.Args[1]
	} else {
		fmt.Print("Enter the documentation website URL to crawl: ")
		fmt.Scanln(&url)
	}

	if url == "" {
		fmt.Println("Error: URL is required")
		return
	}

	fmt.Println("Starting to crawl", url, "...")

	// Create output directories
	dirPath, err := createDirectories(url)
	if err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return
	}

	// Crawl the website
	pages, err := crawlURL(url)
	if err != nil {
		fmt.Printf("Error crawling website: %v\n", err)
		return
	}

	fmt.Printf("Found %d pages\n", len(pages))

	// Write results to files
	if err := writeResults(pages, dirPath); err != nil {
		fmt.Printf("Error writing results: %v\n", err)
		return
	}

	fmt.Println("Crawling completed successfully!")
}
