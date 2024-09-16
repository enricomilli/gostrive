package main

import (
	"flag"
	"fmt"

	"github.com/enricomilli/gostrive/quotes"
)

func main() {
	run()
}

func run() {
	author := flag.String("author", "", "Get a quote from a specific author")
	flag.StringVar(author, "a", "", "Get a quote from a specific author (shorthand)")

	flag.Parse()

	// Check if the author flag is set
	if *author != "" {
		// If author flag is provided, get quotes from that author
		fmt.Println("author:", *author)
		fmt.Println(quotes.FromAuthor(*author))
	} else {
		// If no author flag, get a random quote
		fmt.Println(quotes.Get())
	}
}
