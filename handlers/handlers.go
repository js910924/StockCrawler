package handlers

import (
	"fmt"
	"log"
	"net/http"

	. "StockCrawler/services"
)

func Stock(w http.ResponseWriter, r *http.Request) {
	stockSymbol := r.URL.Query().Get("stockSymbol")
	if stockSymbol == "" {
		fmt.Fprintf(w, "can't find query string: stockSymbol")
		return
	}

	stockCrawler := StockCrawler{
		BaseURL: "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_{stockSymbol}.tw",
	}

	response := stockCrawler.GetStockInfo(stockSymbol)

	if len(response.Stocks) == 0 {
		log.Printf("Can't find stcok: %s\n", stockSymbol)
		fmt.Fprintf(w, "Cant' find stock: %s", stockSymbol)
		return
	}

	fmt.Fprintf(w, "%s", response.Stocks[0].ShowForm())
}

func List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "2324 仁寶\n2317 鴻海\n2330 台積電\n2637 慧洋-KY\n")
}
