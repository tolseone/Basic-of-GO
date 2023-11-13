package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Input URL
	url := "https://google.com" // Replace this with the desired URL

	// Fetch HTML content
	htmlContent, err := fetchHTML(url)
	if err != nil {
		fmt.Println("Error fetching HTML:", err)
		return
	}

	// Write HTML content to a file
	err = writeToFile("output.txt", htmlContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("HTML content written to output.txt")
}

// fetchHTML makes a GET request to the specified URL and retrieves the HTML content.
// It returns the HTML content as a string and an error, if any.
func fetchHTML(url string) (string, error) {
	// Make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	htmlContent := string(htmlBytes)
	return htmlContent, nil
}

// writeToFile writes the provided content to a file with the given filename.
// It returns an error if the operation encounters any issues.
func writeToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}
