package main

import (
	"flag"
	"fmt"
	"log"
	"www.github.com/ah-naf/urlshort/utils"
	"www.github.com/ah-naf/urlshort/models"
)



func main() {
	format := flag.String("config", "yaml", "Choose the configuration format. Options: yaml, json.")
	path := flag.String("path", "redirect.yaml", "Config file path with redirections URLs.")
	flag.Parse()

	var configs []models.Config

	if *format == "yaml" {
		configs = utils.ParseYaml(*path)
	} else if *format == "json" {
		configs = utils.ParseJSON(*path)
	} else {
		log.Fatalln("Invalid config type.")

	}
	fmt.Println(configs)
}
