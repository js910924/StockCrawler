package main

import (
	"fmt"
	"flag"
)

var (
	stockSymbol string
)

func init() {
	flag.StringVar(&stockSymbol, "s", "", "send stock symbol")
}

func main() {
	flag.Parse()

	fmt.Println("Get stock symbol: " + stockSymbol)
}
