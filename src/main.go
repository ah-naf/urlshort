package main

import (
	"flag"
	"fmt"
	"net/http"

	"www.github.com/ah-naf/urlshort/utils"
)

func main() {
	format := flag.String("config", "yaml", "Choose the configuration format. Options: yaml, json.")
	path := flag.String("path", "redirect.yaml", "Config file path with redirections URLs.")
	flag.Parse()

	configs := utils.ParseFile(*format, *path)

	mapPathToUrl := make(map[string]string)
	for _, config := range configs {
		mapPathToUrl[config.Path] = config.URL
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello World!")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if val, ok := mapPathToUrl[path]; ok {
			http.Redirect(w, r, val, http.StatusMovedPermanently)
			return
		}

		http.NotFound(w, r)
	})

	fmt.Println("Server is listening on http://localhost:8080")

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
