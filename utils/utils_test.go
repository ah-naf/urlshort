package utils

import (
	"os"
	"reflect"
	"testing"

	"www.github.com/ah-naf/urlshort/models"
)

func TestParseFile(t *testing.T) {
	// Sample YAML content
	yamlContent := `
- path: "/path/to/resource1"
  url: "https://example1.com"
- path: "/path/to/resource2"
  url: "https://example2.com"
`
	// Sample JSON content
	jsonContent := `[
		{"path": "/path/to/resource1", "url": "https://example1.com"},
		{"path": "/path/to/resource2", "url": "https://example2.com"}
	]`

	// Expected result
	expected := []models.Config{
		{Path: "/path/to/resource1", URL: "https://example1.com"},
		{Path: "/path/to/resource2", URL: "https://example2.com"},
	}

	// Create temporary YAML file
	yamlFile, err := os.CreateTemp("", "test*.yaml")
	if err != nil {
		t.Fatalf("Error creating temporary YAML file: %v", err)
	}
	defer os.Remove(yamlFile.Name()) // Clean up

	_, err = yamlFile.WriteString(yamlContent)
	if err != nil {
		t.Fatalf("Error writing YAML content to file: %v", err)
	}
	yamlFile.Close()

	// Create temporary JSON file
	jsonFile, err := os.CreateTemp("", "test*.json")
	if err != nil {
		t.Fatalf("Error creating temporary JSON file: %v", err)
	}
	defer os.Remove(jsonFile.Name()) // Clean up

	_, err = jsonFile.WriteString(jsonContent)
	if err != nil {
		t.Fatalf("Error writing JSON content to file: %v", err)
	}
	jsonFile.Close()

	// Test YAML parsing
	yamlResult := ParseFile("yaml", yamlFile.Name())
	if !reflect.DeepEqual(yamlResult, expected) {
		t.Errorf("YAML parsing failed. Got %+v, expected %+v", yamlResult, expected)
	}

	// Test JSON parsing
	jsonResult := ParseFile("json", jsonFile.Name())
	if !reflect.DeepEqual(jsonResult, expected) {
		t.Errorf("JSON parsing failed. Got %+v, expected %+v", jsonResult, expected)
	}
}
