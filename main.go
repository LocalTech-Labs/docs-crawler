package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
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

func crawlURL(baseURL string) ([]DocPage, error) {
	// Normalize the input URL
	baseURL = normalizeURL(baseURL)
	fmt.Printf("Starting crawl from: %s\n", baseURL)

	var pages []DocPage
	processedLinks := make(map[string]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Create a channel for new URLs to crawl
	urlQueue := make(chan string, 1000)
	done := make(chan struct{})
	urlQueue <- baseURL

	// Create worker goroutines
	maxWorkers := 5
	fmt.Printf("Starting %d worker goroutines\n", maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case url, ok := <-urlQueue:
					if !ok {
						fmt.Printf("[Worker %d] Channel closed, exiting\n", workerID)
						return
					}

					// Skip if we've already processed this URL
					mu.Lock()
					if processedLinks[url] {
						mu.Unlock()
						continue
					}
					processedLinks[url] = true
					fmt.Printf("[Worker %d] Processing: %s\n", workerID, url)
					mu.Unlock()

					newPages, newLinks, err := crawlPage(url)
					if err != nil {
						fmt.Printf("[Worker %d] Error crawling %s: %v\n", workerID, url, err)
						continue
					}

					mu.Lock()
					pages = append(pages, newPages...)
					fmt.Printf("[Worker %d] Found %d new pages and %d links on %s\n", workerID, len(newPages), len(newLinks), url)
					
					// Add new links to the queue
					for _, link := range newLinks {
						if !processedLinks[link] {
							select {
							case urlQueue <- link:
								fmt.Printf("[Worker %d] Added new URL to queue: %s\n", workerID, link)
							default:
								fmt.Printf("[Worker %d] Queue full, skipping URL: %s\n", workerID, link)
							}
						}
					}
					mu.Unlock()

				case <-done:
					fmt.Printf("[Worker %d] Received done signal, exiting\n", workerID)
					return
				}
			}
		}(i)
	}

	// Start a goroutine to monitor progress
	go func() {
		lastCount := 0
		stuckCount := 0
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				mu.Lock()
				currentCount := len(pages)
				if currentCount == lastCount {
					stuckCount++
					if stuckCount >= 3 { // No progress for 15 seconds
						fmt.Println("No new pages found for 15 seconds, finishing crawl")
						close(done)
						mu.Unlock()
						return
					}
				} else {
					stuckCount = 0
				}
				lastCount = currentCount
				fmt.Printf("Progress update: Found %d pages so far\n", currentCount)
				mu.Unlock()
			}
		}
	}()

	// Wait for all workers to finish
	wg.Wait()
	fmt.Printf("All workers finished. Found total of %d pages\n", len(pages))

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
