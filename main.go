package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/schollz/progressbar/v3"
)

type DocPage struct {
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
}

func createDirectories(baseURL string) (string, error) {
	// Create base directory
	timestamp := time.Now().Format("2006-01-02-15-04-05")
	sanitizedURL := strings.ReplaceAll(strings.ReplaceAll(baseURL, "https://", ""), "/", "-")
	dirPath := filepath.Join("generated_docs", sanitizedURL+"-"+timestamp)
	
	if err := os.MkdirAll(filepath.Join(dirPath, "markdown"), 0755); err != nil {
		return "", err
	}
	
	return dirPath, nil
}

func normalizeURL(url string) string {
	return strings.TrimSuffix(url, "/")
}

func crawlURL(url string) ([]DocPage, error) {
	// Normalize the input URL
	url = normalizeURL(url)

	// Get the webpage
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if it's a 404 page
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("404 page not found")
	}

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var pages []DocPage
	var links []string
	processedLinks := make(map[string]bool)

	// Process the initial page
	title := doc.Find("h1").First().Text()
	if title == "" {
		title = doc.Find("title").First().Text()
	}

	// Find the main content
	var markdown string
	mainContent := doc.Find("article, .markdown, .content, main, #main-content").First()
	if mainContent.Length() > 0 {
		markdown = mainContent.Text()
	}

	if title != "" && markdown != "" && !strings.Contains(title, "Page Not Found") {
		pages = append(pages, DocPage{
			Title:    title,
			Markdown: markdown,
		})
	}

	baseURL := strings.TrimSuffix(url, "/")

	// Find all documentation links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		// Skip external links, anchors, and non-doc pages
		if strings.HasPrefix(href, "http") || strings.HasPrefix(href, "#") {
			return
		}

		// Handle relative URLs
		if strings.HasPrefix(href, "/") {
			href = strings.TrimPrefix(href, "/")
		}

		// Skip non-documentation pages
		if !strings.Contains(href, "docs/") && !strings.HasPrefix(href, "docs") {
			return
		}

		// Construct absolute URL
		absoluteURL := baseURL
		if strings.HasPrefix(href, "docs/") {
			absoluteURL = strings.TrimSuffix(baseURL, "/docs") + "/" + href
		} else if strings.HasPrefix(href, "/docs/") {
			absoluteURL = strings.TrimSuffix(baseURL, "/docs") + href
		} else {
			absoluteURL = baseURL + "/" + strings.TrimPrefix(href, "/")
		}

		// Normalize the absolute URL
		absoluteURL = normalizeURL(absoluteURL)

		// Skip if we've already processed this URL
		if processedLinks[absoluteURL] {
			return
		}
		processedLinks[absoluteURL] = true

		links = append(links, absoluteURL)
	})

	// Create a progress bar
	bar := progressbar.Default(int64(len(links)))

	// Use a WaitGroup to handle concurrent crawling
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Process each link
	for _, link := range links {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			defer bar.Add(1)

			resp, err := http.Get(link)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", link, err)
				return
			}
			defer resp.Body.Close()

			// Skip 404 pages
			if resp.StatusCode == 404 {
				fmt.Printf("Skipping 404 page: %s\n", link)
				return
			}

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				fmt.Printf("Error parsing %s: %v\n", link, err)
				return
			}

			title := doc.Find("h1").First().Text()
			if title == "" {
				title = doc.Find("title").First().Text()
			}

			// Find the main content
			var markdown string
			mainContent := doc.Find("article, .markdown, .content, main, #main-content").First()
			if mainContent.Length() > 0 {
				markdown = mainContent.Text()
			}

			if title != "" && markdown != "" {
				page := DocPage{
					Title:    title,
					Markdown: markdown,
				}
				mu.Lock()
				pages = append(pages, page)
				mu.Unlock()
			}
		}(link)
	}

	wg.Wait()
	return pages, nil
}

func writeResults(pages []DocPage, dirPath string) error {
	// Write the crawl results
	resultsFile, err := os.Create(filepath.Join(dirPath, "crawl_results.json"))
	if err != nil {
		return err
	}
	defer resultsFile.Close()

	encoder := json.NewEncoder(resultsFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(pages); err != nil {
		return err
	}

	// Write individual markdown files
	for _, page := range pages {
		// Sanitize the filename
		filename := strings.Map(func(r rune) rune {
			switch {
			case r >= 'a' && r <= 'z':
				return r
			case r >= 'A' && r <= 'Z':
				return r
			case r >= '0' && r <= '9':
				return r
			case r == ' ' || r == '-' || r == '_' || r == '.':
				return r
			default:
				return '_'
			}
		}, page.Title)

		filename = strings.ToLower(filename) + ".md"
		
		mdFile, err := os.Create(filepath.Join(dirPath, "markdown", filename))
		if err != nil {
			return err
		}
		defer mdFile.Close()

		writer := bufio.NewWriter(mdFile)
		_, err = writer.WriteString(page.Markdown)
		if err != nil {
			return err
		}
		writer.Flush()
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the documentation website URL to crawl: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	// Create directories
	dirPath, err := createDirectories(url)
	if err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return
	}

	fmt.Printf("Starting to crawl %s...\n", url)
	pages, err := crawlURL(url)
	if err != nil {
		fmt.Printf("Error crawling website: %v\n", err)
		return
	}

	fmt.Printf("Found %d pages. Writing results...\n", len(pages))
	if err := writeResults(pages, dirPath); err != nil {
		fmt.Printf("Error writing results: %v\n", err)
		return
	}

	fmt.Printf("Crawling complete! Results saved in %s\n", dirPath)
}
