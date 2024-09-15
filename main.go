package main

import (
	"fmt"

	"github.com/enricomilli/gostrive/quotes"
)

func main() {
	run()
}

func run() {
	fmt.Println(quotes.Get())
}
