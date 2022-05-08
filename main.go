package main

import (
	"flag"
	"fmt"
	goiplookup "github.com/heavykenny/goiplookup/src"
)

var verbose bool
var url, path, output string

func main() {
	flag.StringVar(&url, "url", "", "url")
	flag.StringVar(&url, "u", "", "url")

	flag.StringVar(&path, "p", "", "path")
	flag.StringVar(&path, "path", "", "path")

	flag.StringVar(&output, "o", "", "output")
	flag.StringVar(&output, "output", "", "output")

	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.BoolVar(&verbose, "verbose", false, "verbose")

	flag.Usage = func() { fmt.Print(goiplookup.Usage) }
	flag.Parse()

	if url != "" {
		goiplookup.GetIPFromURL(goiplookup.FilterURL(url), verbose)
	}

	if path != "" {
		goiplookup.ReadFile(path, verbose, output)
	}
}
