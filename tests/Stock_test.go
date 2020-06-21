package testing

import (
	. "StockCrawler/models"
	"testing"
)

func TestWhenStockShowForm(t *testing.T) {
	stock := givenStock()
	actual := stock.ShowForm()
	expected := `StockSymbol: 2618
StockName: 長榮航
CurrencyTransactionPrice: 12
CurrentVolumn: 2000
OpeningPrice: 12
HighestPrice: 12
YesturdayClosePrice: 12
LowestPrice: 12
Percent: 0.000%
`

	if actual != expected {
		t.Errorf("\n%s is not equal\n%s", actual, expected)
	}
}

func givenStock() Stock {
	return Stock{
		StockSymbol:              "2618",
		StockName:                "長榮航",
		CurrencyTransactionPrice: "12",
		CurrentVolumn:            "2000",
		OpeningPrice:             "12",
		HighestPrice:             "12",
		YesturdayClosePrice:      "12",
		LowestPrice:              "12",
	}
}
