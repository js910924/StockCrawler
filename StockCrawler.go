package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type StockCrawler struct {
	baseURL string
}

func New(baseURL string) StockCrawler {
	return StockCrawler{
		baseURL: baseURL,
	}
}

func (s *StockCrawler) GetStockInfo(stockSymbol string) GetStockInfoResponse {
	url := strings.Replace(s.baseURL, "{stockSymbol}", stockSymbol, 1)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("No such stock symbol: "+stockSymbol, ", Error: ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Read response body error")
	}

	defer response.Body.Close()

	var getStockInfoResponse GetStockInfoResponse
	json.Unmarshal([]byte(body), &getStockInfoResponse)

	return getStockInfoResponse
}

type GetStockInfoResponse struct {
	// CachedAlive int64  `json:"cachedAlive"`
	// ExKey       string `json:"exKey"`
	Stocks []Stock `json:"msgArray"`
	// QueryTime struct {
	// 	SessionFromTime   int64  `json:"sessionFromTime"`
	// 	SessionKey        string `json:"sessionKey"`
	// 	SessionLatestTime int64  `json:"sessionLatestTime"`
	// 	SessionStr        string `json:"sessionStr"`
	// 	ShowChart         bool   `json:"showChart"`
	// 	StockInfo         int64  `json:"stockInfo"`
	// 	StockInfoItem     int64  `json:"stockInfoItem"`
	// 	SysDate           string `json:"sysDate"`
	// 	SysTime           string `json:"sysTime"`
	// } `json:"queryTime"`
	// Referer   string `json:"referer"`
	// Rtcode    string `json:"rtcode"`
	Rtmessage string `json:"rtmessage"`
	// UserDelay int64  `json:"userDelay"`
}

type Stock struct {
	// A                        string `json:"a"`
	// B                        string `json:"b"`
	StockSymbol              string `json:"c"`
	StockName                string `json:"n"`
	CurrencyTransactionPrice string `json:"z"`
	CurrentVolumn            string `json:"tv"`
	OpeningPrice             string `json:"o"`
	HighestPrice             string `json:"h"`
	YesturdayClosePrice      string `json:"y"`
	LowestPrice              string `json:"l"`
	// Ch                       string `json:"ch"`
	// D                        string `json:"d"`
	// Ex                       string `json:"ex"`
	// F                        string `json:"f"`
	// Fv                       string `json:"fv"`
	// G                        string `json:"g"`
	// I                        string `json:"i"`
	// IP                       string `json:"ip"`
	// It                       string `json:"it"`
	// Mt                       string `json:"mt"`
	// Nf                       string `json:"nf"`
	// Oa                       string `json:"oa"`
	// Ob                       string `json:"ob"`
	// Ot                       string `json:"ot"`
	// Ov                       string `json:"ov"`
	// Oz                       string `json:"oz"`
	// P                        string `json:"p"`
	// Ps                       string `json:"ps"`
	// Pz                       string `json:"pz"`
	// S                        string `json:"s"`
	// T                        string `json:"t"`
	// Tk0                      string `json:"tk0"`
	// Tk1                      string `json:"tk1"`
	// Tlong                    string `json:"tlong"`
	// Ts                       string `json:"ts"`
	// U                        string `json:"u"`
	// W                        string `json:"w"`
}
