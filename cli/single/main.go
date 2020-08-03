package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "first name", "help")
	flag.StringVar(&name, "n", "first name", "help")
	flag.Parse()
	log.Printf("name %s", name)
}
