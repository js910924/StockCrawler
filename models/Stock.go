package models

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

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

func (s *Stock) ShowForm() {
	name := reflect.Indirect(reflect.ValueOf(s))
	for i := 0; i < name.NumField(); i++ {
		fmt.Printf("%s %s\n", name.Type().Field(i).Name, name.Field(i).Interface())
	}

	fmt.Printf("Percent %.3f%%\n", s.GetPercent())
}

func (s *Stock) GetPercent() float64 {
	currenctPrice, err := strconv.ParseFloat(s.CurrencyTransactionPrice, 0)
	if err != nil {
		log.Fatalln("Parse fail")
	}

	openingPrice, err := strconv.ParseFloat(s.OpeningPrice, 0)
	if err != nil {
		log.Fatalln("Parse fail")
	}

	return (currenctPrice - openingPrice) / openingPrice * 100
}
