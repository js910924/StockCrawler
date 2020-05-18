package main

import (
	"fmt"
	"flag"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"encoding/json"
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
	url := "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_{stockSymbol}.tw"
	url = strings.Replace(url, "{stockSymbol}", stockSymbol, 1)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("No such stock symbol: " + stockSymbol, ", Error: ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Read response body error")
	}

	defer response.Body.Close()

	var getStockInfoResponse GetStockInfoResponse
	json.Unmarshal([]byte(body), &getStockInfoResponse)

	fmt.Printf("%+v\n", getStockInfoResponse.MsgArray[0])
}

type GetStockInfoResponse struct {
	CachedAlive int64  `json:"cachedAlive"`
	ExKey       string `json:"exKey"`
	MsgArray    []struct {
		A     string `json:"a"`
		B     string `json:"b"`
		C     string `json:"c"`
		Ch    string `json:"ch"`
		D     string `json:"d"`
		Ex    string `json:"ex"`
		F     string `json:"f"`
		Fv    string `json:"fv"`
		G     string `json:"g"`
		H     string `json:"h"`
		I     string `json:"i"`
		IP    string `json:"ip"`
		It    string `json:"it"`
		L     string `json:"l"`
		Mt    string `json:"mt"`
		N     string `json:"n"`
		Nf    string `json:"nf"`
		O     string `json:"o"`
		Oa    string `json:"oa"`
		Ob    string `json:"ob"`
		Ot    string `json:"ot"`
		Ov    string `json:"ov"`
		Oz    string `json:"oz"`
		P     string `json:"p"`
		Ps    string `json:"ps"`
		Pz    string `json:"pz"`
		S     string `json:"s"`
		T     string `json:"t"`
		Tk0   string `json:"tk0"`
		Tk1   string `json:"tk1"`
		Tlong string `json:"tlong"`
		Ts    string `json:"ts"`
		Tv    string `json:"tv"`
		U     string `json:"u"`
		V     string `json:"v"`
		W     string `json:"w"`
		Y     string `json:"y"`
		Z     string `json:"z"`
	} `json:"msgArray"`
	QueryTime struct {
		SessionFromTime   int64  `json:"sessionFromTime"`
		SessionKey        string `json:"sessionKey"`
		SessionLatestTime int64  `json:"sessionLatestTime"`
		SessionStr        string `json:"sessionStr"`
		ShowChart         bool   `json:"showChart"`
		StockInfo         int64  `json:"stockInfo"`
		StockInfoItem     int64  `json:"stockInfoItem"`
		SysDate           string `json:"sysDate"`
		SysTime           string `json:"sysTime"`
	} `json:"queryTime"`
	Referer   string `json:"referer"`
	Rtcode    string `json:"rtcode"`
	Rtmessage string `json:"rtmessage"`
	UserDelay int64  `json:"userDelay"`
}