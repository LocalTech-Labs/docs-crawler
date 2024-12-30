package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteResults(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test-docs-*")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir) // Clean up after test

	// Create the markdown subdirectory
	markdownDir := filepath.Join(tempDir, "markdown")
	err = os.MkdirAll(markdownDir, 0755)
	assert.NoError(t, err)

	// Test cases
	tests := []struct {
		name     string
		pages    []DocPage
		wantErr  bool
		validate func(t *testing.T, dir string)
	}{
		{
			name: "Single page",
			pages: []DocPage{
				{
					Title:    "Test Page",
					Markdown: "# Test Content\nThis is a test.",
				},
			},
			wantErr: false,
			validate: func(t *testing.T, dir string) {
				// Check if file exists
				path := filepath.Join(dir, "markdown", "test-page.md")
				_, err := os.Stat(path)
				assert.NoError(t, err)

				// Check file contents
				content, err := os.ReadFile(path)
				assert.NoError(t, err)
				assert.Contains(t, string(content), "# Test Content")
				assert.Contains(t, string(content), "This is a test.")
			},
		},
		{
			name:    "Empty pages",
			pages:   []DocPage{},
			wantErr: false,
			validate: func(t *testing.T, dir string) {
				files, err := os.ReadDir(filepath.Join(dir, "markdown"))
				assert.NoError(t, err)
				assert.Empty(t, files)
			},
		},
		{
			name: "Multiple pages",
			pages: []DocPage{
				{
					Title:    "Page One",
					Markdown: "# Page One Content",
				},
				{
					Title:    "Page Two",
					Markdown: "# Page Two Content",
				},
			},
			wantErr: false,
			validate: func(t *testing.T, dir string) {
				files, err := os.ReadDir(filepath.Join(dir, "markdown"))
				assert.NoError(t, err)
				assert.Len(t, files, 2)
			},
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean the markdown directory before each test
			err := os.RemoveAll(markdownDir)
			assert.NoError(t, err)
			err = os.MkdirAll(markdownDir, 0755)
			assert.NoError(t, err)

			// Run the function
			err = writeResults(tt.pages, tempDir)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				tt.validate(t, tempDir)
			}
		})
	}
}

func TestCreateDirectories(t *testing.T) {
	// Clean up any existing test directories
	defer os.RemoveAll("generated_docs")

	tests := []struct {
		name     string
		baseURL  string
		wantErr  bool
		validate func(t *testing.T, dirPath string)
	}{
		{
			name:    "Simple URL",
			baseURL: "https://example.com",
			wantErr: false,
			validate: func(t *testing.T, dirPath string) {
				// Check main directory exists
				assert.DirExists(t, dirPath)
				// Check markdown subdirectory exists
				assert.DirExists(t, filepath.Join(dirPath, "markdown"))
				// Verify directory name
				assert.Equal(t, "generated_docs/example.com", dirPath)
			},
		},
		{
			name:    "URL with path",
			baseURL: "https://example.com/docs/api",
			wantErr: false,
			validate: func(t *testing.T, dirPath string) {
				assert.DirExists(t, dirPath)
				assert.DirExists(t, filepath.Join(dirPath, "markdown"))
				assert.Equal(t, "generated_docs/example.com-docs-api", dirPath)
			},
		},
		{
			name:    "URL with special characters",
			baseURL: "https://example.com/docs?version=1.0&lang=en",
			wantErr: false,
			validate: func(t *testing.T, dirPath string) {
				assert.DirExists(t, dirPath)
				assert.DirExists(t, filepath.Join(dirPath, "markdown"))
				assert.Equal(t, "generated_docs/example.com-docs?version=1.0&lang=en", dirPath)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up before each test
			os.RemoveAll("generated_docs")

			// Run the function
			dirPath, err := createDirectories(tt.baseURL)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			tt.validate(t, dirPath)
		})
	}
}
