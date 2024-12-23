package main

import (
	"flag"
	"fmt"

	"www.github.com/ah-naf/urlshort/utils"
)

func main() {
	format := flag.String("config", "yaml", "Choose the configuration format. Options: yaml, json.")
	path := flag.String("path", "redirect.yaml", "Config file path with redirections URLs.")
	flag.Parse()

	configs := utils.ParseFile(*format, *path)

	fmt.Println(configs)
}
