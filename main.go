package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	. "StockCrawler/services"
	. "StockCrawler/tools"
)

var (
	stockSymbol string
	isShowList  bool
)

func init() {
	flag.StringVar(&stockSymbol, "s", "", "send stock symbol")
	flag.BoolVar(&isShowList, "l", false, "show favorite stocks")
}

func main() {
	flag.Parse()
	stockCrawler := StockCrawler{
		BaseURL: "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_{stockSymbol}.tw",
	}

	if stockSymbol != "" {
		response := stockCrawler.GetStockInfo(stockSymbol)
		fmt.Printf("RtMessage: %s\n%s", response.Rtmessage, response.Stocks[0].ShowForm())
		return
	}

	if isShowList {
		GetFavoriteStockList()
		return
	}

	http.HandleFunc("/Stock", func(w http.ResponseWriter, r *http.Request) {
		stockSymbol = r.URL.Query().Get("stockSymbol")
		if stockSymbol == "" {
			fmt.Fprintf(w, "can't find query string: stockSymbol")
			return
		}

		response := stockCrawler.GetStockInfo(stockSymbol)

		if len(response.Stocks) == 0 {
			log.Printf("Can't find stcok: %s\n", stockSymbol)
			fmt.Fprintf(w, "Cant' find stock: %s", stockSymbol)
			return
		}

		fmt.Fprintf(w, "%s", response.Stocks[0].ShowForm())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
