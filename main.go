package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {

	fmt.Println(req)

	url := req.RequestURI[1:]

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
	fmt.Fprintf(w, string(htmlContent))

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
	return os.WriteFile("shared-data/"+filename, content, 0644)
}
