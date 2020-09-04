package main

import (
	"net/http"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestBittrexAPIFixture(t *testing.T) {
	gunit.Run(new(BittrexAPIFixture), t)

}

type BittrexAPIFixture struct {
	*gunit.Fixture
}

func (this *BittrexAPIFixture) Setup() {}

func (this *BittrexAPIFixture) TestGetCurrency() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "fakeURL")
	result := bittrex.getCurrency("btc")
	this.So(result, should.Resemble, Currency{})
}

func (this *BittrexAPIFixture) TestGetBalances() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result := bittrex.getBalances()
	this.So(result, should.Resemble, []Balance{
		{
			CurrencySymbol: "BTC",
			Total:          "0.00000000",
			Available:      "0.00000000",
			UpdatedAt:      "2019-10-29T20:25:10.16Z",
		}, {
			CurrencySymbol: "LTC",
			Total:          "0",
			Available:      "0",
			UpdatedAt:      "2020-09-03T21:27:53.8210894Z",
		},
	})
}
func (this *BittrexAPIFixture) TestGetMarket() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result := bittrex.getMarket("fakesymbol")
	this.So(result, should.Resemble, Market{
		Symbol:              "ETH-BTC",
		BaseCurrencySymbol:  "ETH",
		QuoteCurrencySymbol: "BTC",
		MinTradeSize:        "0.01000000",
		Precision:           8,
		Status:              "ONLINE",
		CreatedAt:           "2015-08-14T09:02:24.817Z",
		Notice:              "",
		ProhibitedIn:        []string{},
	})
}

func (this *BittrexAPIFixture) TestGetMarkets() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result := bittrex.getMarkets()
	this.So(result, should.Resemble, []Market{
		{
			Symbol:              "4ART-BTC",
			BaseCurrencySymbol:  "4ART",
			QuoteCurrencySymbol: "BTC",
			MinTradeSize:        "10.00000000",
			Precision:           8,
			Status:              "ONLINE",
			CreatedAt:           "2020-06-10T15:05:29.833Z",
			Notice:              "",
			ProhibitedIn:        []string{"US"},
		}, {
			Symbol:              "4ART-USDT",
			BaseCurrencySymbol:  "4ART",
			QuoteCurrencySymbol: "USDT",
			MinTradeSize:        "10.00000000",
			Precision:           5,
			Status:              "ONLINE",
			CreatedAt:           "2020-06-10T15:05:40.98Z",
			Notice:              "",
			ProhibitedIn:        []string{"US"},
		},
	})
}

///////////////////////////////////////

type fakeBittrexClient struct{}

func (this *fakeBittrexClient) Do(method, uri, payload string, authenticate bool) []byte {
	if uri == "/balances" {
		return []byte("[{\"currencySymbol\": \"BTC\",\"total\": \"0.00000000\",\"available\": \"0.00000000\",\"updatedAt\": \"2019-10-29T20:25:10.16Z\"},{\"currencySymbol\": \"LTC\",\"total\": \"0\",\"available\": \"0\",\"updatedAt\": \"2020-09-03T21:27:53.8210894Z\"}]")
	}

	if uri == "/markets/fakesymbol" {
		return []byte("{\"symbol\":\"ETH-BTC\",\"baseCurrencySymbol\":\"ETH\",\"quoteCurrencySymbol\":\"BTC\",\"minTradeSize\":\"0.01000000\",\"precision\":8,\"status\":\"ONLINE\",\"createdAt\":\"2015-08-14T09:02:24.817Z\",\"notice\":\"\",\"prohibitedIn\":[],\"associatedTermsOfService\":[]}")
	}

	if uri == "/markets" {
		return []byte("[{\"symbol\": \"4ART-BTC\",\"baseCurrencySymbol\": \"4ART\",\"quoteCurrencySymbol\": \"BTC\",\"minTradeSize\": \"10.00000000\",\"precision\": 8,\"status\": \"ONLINE\",\"createdAt\": \"2020-06-10T15:05:29.833Z\",\"notice\": \"\",\"prohibitedIn\": [\"US\"]},{\"symbol\": \"4ART-USDT\",\"baseCurrencySymbol\": \"4ART\",\"quoteCurrencySymbol\": \"USDT\",\"minTradeSize\": \"10.00000000\",\"precision\": 5,\"status\": \"ONLINE\",\"createdAt\": \"2020-06-10T15:05:40.98Z\",\"notice\": \"\",\"prohibitedIn\": [\"US\"]}]")
	}
	return nil
}

func (this *fakeBittrexClient) authenticate(request *http.Request, payload string, uri string, method string) error {
	return nil
}
