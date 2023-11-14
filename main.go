package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	url := "google.com"
	// var url string
	// fmt.Print("Enter the URL: ")
	// fmt.Scan(&url)

	if !(strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")) {
		url = "https://" + url
	}

	htmlContent, err := fetchHTML(url)
	if err != nil {
		fmt.Println("Error fetching HTML:", err)
		return
	}

	err = writeToFile("output.txt", htmlContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("HTML content written to output.txt")

	http.HandleFunc("/docker", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, string(htmlContent))
	})
	http.ListenAndServe(":8080", nil)
}

func fetchHTML(url string) ([]byte, error) {
	// Make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return htmlBytes, nil
}

func writeToFile(filename string, content []byte) error {
	return os.WriteFile(filename, content, 0644)
}
