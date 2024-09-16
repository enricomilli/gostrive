package main

import (
	"fmt"

	"github.com/enricomilli/gostrive/quotes"
)

func main() {
	run()
}

func run() {

	// if not flags
	fmt.Println(quotes.Get())

	// if author flag
	// quotes.FromAuthor(author)
}
