package main

import (
	"flag"
	"fmt"
)

var (
	stockSymbol string
)

func init() {
	flag.StringVar(&stockSymbol, "s", "", "send stock symbol")
}

func main() {
	flag.Parse()

	url := "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_{stockSymbol}.tw"
	stockCrawler := StockCrawler{
		baseURL: url,
	}
	fmt.Println(stockCrawler.GetStockInfo(stockSymbol))
}
