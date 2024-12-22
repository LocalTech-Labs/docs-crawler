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
		content, err := extractContent(mainContent)
		if err != nil {
			return nil, fmt.Errorf("error extracting content: %v", err)
		}
		markdown = formatMarkdown(title, content)
	}

	if title != "" && markdown != "" && !strings.Contains(title, "Page Not Found") {
		pages = append(pages, DocPage{
			Title:    title,
			Markdown: markdown,
		})
	}

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
		absoluteURL := url
		if strings.HasPrefix(href, "docs/") {
			absoluteURL = strings.TrimSuffix(url, "/docs") + "/" + href
		} else {
			absoluteURL = strings.TrimSuffix(url, "/docs") + "/docs/" + strings.TrimPrefix(href, "docs/")
		}

		// Skip if we've already processed this URL
		if processedLinks[absoluteURL] {
			return
		}
		processedLinks[absoluteURL] = true

		links = append(links, absoluteURL)
	})

	var wg sync.WaitGroup
	var mu sync.Mutex

	// Process each link
	for _, link := range links {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()

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
				content, err := extractContent(mainContent)
				if err != nil {
					fmt.Printf("Error extracting content from %s: %v\n", link, err)
					return
				}
				markdown = formatMarkdown(title, content)
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
		os.Exit(1)
	}

	fmt.Println("Starting to crawl ...")
	pages, err := crawlURL(url)
	if err != nil {
		fmt.Printf("Error crawling website: %v\n", err)
		os.Exit(1)
	}

	// Create output directory
	outputDir := filepath.Join(".", "generated_docs", strings.TrimPrefix(strings.TrimPrefix(url, "https://"), "http://"))
	if err := os.MkdirAll(filepath.Join(outputDir, "markdown"), 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Write results
	if err := writeResults(pages, outputDir); err != nil {
		fmt.Printf("Error writing results: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully crawled %d pages. Results written to %s\n", len(pages), outputDir)
}
