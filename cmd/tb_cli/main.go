package main

import (
	"fmt"

	"github.com/piquette/finance-go/quote"
)

func main() {
	q, _ := quote.Get("AAPL")
	fmt.Println(q)
}
