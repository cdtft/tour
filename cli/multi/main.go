package main

import (
	"flag"
	"log"
)

var name string

func main() {
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "", "help")

	javaCmd := flag.NewFlagSet("java", flag.ExitOnError)
	javaCmd.StringVar(&name, "name", "", "help")

	args := flag.Args()
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "java":
		_ = javaCmd.Parse(args[1:])
	}

	log.Printf("name %s", name)
}
