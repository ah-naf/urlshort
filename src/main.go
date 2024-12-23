package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"www.github.com/ah-naf/urlshort/utils"
)

func isValidURL(u string) bool {
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}

	validSchemes := map[string]bool{
		"http":   true,
		"https":  true,
		"ftp":    true,
		"ws":     true,
		"wss":    true,
		"mailto": true,
		"tel":    true,
		"data":   true,
		"sftp":   true,
	}

	if !validSchemes[parsedURL.Scheme] {
		return false
	}

	host := parsedURL.Host
	if strings.Contains(host, "..") || strings.ContainsAny(host, " !@#$%^&*(){}[]|\\<>") {
		return false
	}

	return true
}

func main() {
	format := flag.String("config", "yaml", "Choose the configuration format. Options: yaml, json.")
	path := flag.String("path", "redirect.yaml", "Config file path with redirections URLs.")
	flag.Parse()

	configs := utils.ParseFile(*format, *path)
	mapPathToUrl := make(map[string]string)

	totalURLs := len(configs)
	validURLs := 0
	invalidURLs := 0
	for _, config := range configs {
		if isValidURL(config.URL) {
			mapPathToUrl[config.Path] = config.URL
			validURLs++
		} else {
			invalidURLs++
		}
	}

	log.Printf("Total URLs: %d\n", totalURLs)
	log.Printf("Valid URLs: %d\n", validURLs)
	log.Printf("Invalid URLs: %d\n", invalidURLs)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello World!")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if val, ok := mapPathToUrl[path]; ok {
			log.Printf("Redirecting %s to %s\n", path, val)
			http.Redirect(w, r, val, http.StatusMovedPermanently)
			return
		}

		log.Printf("Path not found: %s\n", path)
		http.NotFound(w, r)
	})

	log.Println("Server is listening on http://localhost:8080")

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Error starting server: %v\n", err)
	}
}
