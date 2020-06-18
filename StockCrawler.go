package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	. "StockCrawler/models"
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
