package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

func main() {
	format := flag.String("config", "yaml", "Choose the configuration format. Options: yaml, json.")
	path := flag.String("path", "redirect.yaml", "Config file path with redirections URLs.")
	flag.Parse()

	if *format == "yaml" {
		parseYaml(*path)
	} else if *format == "json" {

	} else {
		log.Fatalln("Invalid config type.")

	}
}

func parseYaml(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer file.Close()

	var configs []Config

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&configs)
	if err != nil {
		log.Fatalf("Error decoding YAML: %v\n", err)
	}

	// Output the parsed data
	for i, config := range configs {
		fmt.Printf("Config %d: %+v\n", i+1, config)
	}
}
