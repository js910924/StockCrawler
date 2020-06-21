package testing

import (
	. "StockCrawler/models"
	"strings"
	"testing"
)

func TestWhenStockShowForm(t *testing.T) {
	stock := Stock{
		StockSymbol:              "2618",
		StockName:                "長榮航",
		CurrencyTransactionPrice: "12",
		CurrentVolumn:            "2000",
		OpeningPrice:             "12",
		HighestPrice:             "12",
		YesturdayClosePrice:      "12",
		LowestPrice:              "12",
	}

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

	if strings.Compare(actual, expected) != 0 {
		t.Errorf("\n%s is not equal\n%s", stock.ShowForm(), expected)
	}
}
