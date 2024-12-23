package utils

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
	"www.github.com/ah-naf/urlshort/models"
)

func ParseFile(format, filePath string) []models.Config {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer file.Close()

	var configs []models.Config

	format = strings.ToLower(format)
	if format == "yaml" {
		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(&configs)
	} else if format == "json" {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&configs)
	} else {
		log.Fatalln("Invalid config type.")
	}

	if err != nil {
		log.Fatalf("Error decoding %s: %v\n", strings.ToUpper(format), err)
	}

	return configs
}
