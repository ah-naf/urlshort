package utils

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"www.github.com/ah-naf/urlshort/models"
)

func ParseYaml(filePath string) []models.Config {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer file.Close()

	var configs []models.Config

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&configs)
	if err != nil {
		log.Fatalf("Error decoding YAML: %v\n", err)
	}

	// Output the parsed data
	// for i, config := range configs {
	// 	fmt.Printf("Config %d: %+v\n", i+1, config)
	// }
	return configs
}

func ParseJSON(filePath string) []models.Config {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer file.Close()

	var configs []models.Config

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configs)

	if err != nil {
		log.Fatalf("Error decoding JSON: %v\n", err)
	}

	return configs
}
