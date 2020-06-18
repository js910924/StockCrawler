package main

import (
	"flag"
	"fmt"

	. "StockCrawler/services"
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
		BaseURL: url,
	}

	response := stockCrawler.GetStockInfo(stockSymbol)
	fmt.Println(response.Rtmessage)
	fmt.Println(response.Stocks[0])
	response.Stocks[0].ShowForm()
}
