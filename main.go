package main

import (
	"fmt"
	"flag"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
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
	GetStockPicFromYahoo(stockSymbol)
}

func GetStockPicFromYahoo(stockSymbol string) {
	url := "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_{stockSymbol}.tw&json=1&delay=0&_=1552123547443"
	url = strings.Replace(url, "{stockSymbol}", stockSymbol, 1)
	response, err := http.Get(url + stockSymbol)
	if err != nil {
		log.Fatalln("No such stock symbol: " + stockSymbol)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Read response body error")
	}

	defer response.Body.Close()

	fmt.Printf("%s\n", body)
}
