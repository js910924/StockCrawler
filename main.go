package main

import (
	"flag"
	"log"
	"net/http"

	"StockCrawler/handlers"
	. "StockCrawler/tools"
)

var (
	isShowList bool
)

func init() {
	flag.BoolVar(&isShowList, "l", false, "show favorite stocks")
}

func main() {
	flag.Parse()

	if isShowList {
		GetFavoriteStockList()
		return
	}

	r := http.NewServeMux()

	r.HandleFunc("/Stock", handlers.Stock)

	log.Fatal(http.ListenAndServe(":8080", r))
}
