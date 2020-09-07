package bittrex

import (
	"errors"
	"net/http"
	"testing"

	"github.com/shopspring/decimal"
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
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.getCurrency("fakesymbol")
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, Currency{})
}

func (this *BittrexAPIFixture) TestGetBalances() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.getBalances()
	this.So(err, should.BeNil)
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
	result, err := bittrex.getMarket("fakesymbol")
	this.So(err, should.BeNil)
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
	result, err := bittrex.getMarkets()
	this.So(err, should.BeNil)
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

func (this *BittrexAPIFixture) TestGetMarketSummary() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.getMarketSummary("fakesymbol")
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, MarketSummary{
		Symbol:        "ETH-BTC",
		High:          decimal.NewFromFloatWithExponent(0.03894964, -8),
		Low:           decimal.NewFromFloatWithExponent(0.03650000, -8),
		Volume:        decimal.NewFromFloat(18494.04035144),
		QuoteVolume:   decimal.NewFromFloat(696.42899671),
		PercentChange: decimal.NewFromFloat(-3.33),
		UpdatedAt:     "2020-09-04T04:37:45.107Z",
	})
}

func (this *BittrexAPIFixture) TestGetMarketSummaries() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.getMarketSummaries()
	this.So(err, should.BeNil)

	this.So(result, should.Resemble, []MarketSummary{
		{
			Symbol:        "4ART-BTC",
			High:          decimal.NewFromFloat(0.00000275),
			Low:           decimal.NewFromFloat(0.00000249),
			Volume:        decimal.NewFromFloat(54499.59344453),
			QuoteVolume:   decimal.NewFromFloat(0.13917073),
			PercentChange: decimal.NewFromFloat(10.44),
			UpdatedAt:     "2020-09-04T04:58:55.447Z",
		}, {
			Symbol:        "4ART-USDT",
			High:          decimal.NewFromFloatWithExponent(0.02880000, -8),
			Low:           decimal.NewFromFloatWithExponent(0.02667000, -8),
			Volume:        decimal.NewFromFloat(48259.53706735),
			QuoteVolume:   decimal.NewFromFloat(1320.75839607),
			PercentChange: decimal.NewFromFloat(-6.11),
			UpdatedAt:     "2020-09-04T04:33:20.01Z",
		},
	})
}

func (this *BittrexAPIFixture) TestGetMarketTicker() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.getMarketTicker("fakesymbol")
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, MarketTicker{
		Symbol:        "ETH-BTC",
		LastTradeRate: decimal.NewFromFloat(0.03760069),
		BidRate:       decimal.NewFromFloat(0.03760103),
		AskRate:       decimal.NewFromFloat(0.03762798),
	})
}

func (this *BittrexAPIFixture) TestGetMarketTickers() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.getMarketTickers()
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, []MarketTicker{
		{
			Symbol:        "ETH-BTC",
			LastTradeRate: decimal.NewFromFloat(0.03760069),
			BidRate:       decimal.NewFromFloat(0.03760103),
			AskRate:       decimal.NewFromFloat(0.03762798),
		},
		{
			Symbol:        "ETH-FAKE",
			LastTradeRate: decimal.NewFromFloat(1.03760069),
			BidRate:       decimal.NewFromFloat(1.03760103),
			AskRate:       decimal.NewFromFloat(1.03762798),
		},
	})
}

func (this *BittrexAPIFixture) TestGetOrder() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.getOrder("fakeOrder")
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, Order{
		OrderID:      "55eb2c82-4184-4a24-8b6e-ee154b2f7eaf",
		MarketSymbol: "XRP-BTC",
		Direction:    "BUY",
		OrderType:    "LIMIT",
		Quantity:     decimal.NewFromFloatWithExponent(77.53046131, -8),
		Limit:        decimal.NewFromFloatWithExponent(0.00003528, -8),
		TimeInForce:  "GOOD_TIL_CANCELLED",
		FillQuantity: decimal.NewFromFloatWithExponent(77.53046131, -8),
		Commission:   decimal.NewFromFloatWithExponent(0.00000682, -8),
		Proceeds:     decimal.NewFromFloatWithExponent(0.00272829, -8),
		Status:       "CLOSED",
		CreatedAt:    "2017-10-20T18:27:20.747Z",
		UpdatedAt:    "2017-10-20T18:27:20.763Z",
		ClosedAt:     "2017-10-20T18:27:20.763Z",
	})
}
func (this *BittrexAPIFixture) TestGetOrders() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.getOrders("open")
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, []Order{{
		OrderID:      "55eb2c82-4184-4a24-8b6e-ee154b2f7eaf",
		MarketSymbol: "XRP-BTC",
		Direction:    "BUY",
		OrderType:    "LIMIT",
		Quantity:     decimal.NewFromFloatWithExponent(77.53046131, -8),
		Limit:        decimal.NewFromFloatWithExponent(0.00003528, -8),
		TimeInForce:  "GOOD_TIL_CANCELLED",
		FillQuantity: decimal.NewFromFloatWithExponent(77.53046131, -8),
		Commission:   decimal.NewFromFloatWithExponent(0.00000682, -8),
		Proceeds:     decimal.NewFromFloatWithExponent(0.00272829, -8),
		Status:       "OPEN",
		CreatedAt:    "2017-10-20T18:27:20.747Z",
		UpdatedAt:    "2017-10-20T18:27:20.763Z",
		ClosedAt:     "2017-10-20T18:27:20.763Z",
	}})

	result, err = bittrex.getOrders("closed")
	this.So(err, should.BeNil)
	this.So(result[0].Status, should.Equal, "CLOSED")
}

///////////////////////////////////////

type fakeBittrexClient struct{}

func (this *fakeBittrexClient) Do(method, uri, payload string, authenticate bool) ([]byte, error) {
	if uri == "/balances" {
		return []byte("[{\"currencySymbol\": \"BTC\",\"total\": \"0.00000000\",\"available\": \"0.00000000\",\"updatedAt\": \"2019-10-29T20:25:10.16Z\"},{\"currencySymbol\": \"LTC\",\"total\": \"0\",\"available\": \"0\",\"updatedAt\": \"2020-09-03T21:27:53.8210894Z\"}]"), nil
	}

	if uri == "/currencies/fakesymbol" {
		return []byte("{}"), nil
	}

	if uri == "/markets/fakesymbol" {
		return []byte("{\"symbol\":\"ETH-BTC\",\"baseCurrencySymbol\":\"ETH\",\"quoteCurrencySymbol\":\"BTC\",\"minTradeSize\":\"0.01000000\",\"precision\":8,\"status\":\"ONLINE\",\"createdAt\":\"2015-08-14T09:02:24.817Z\",\"notice\":\"\",\"prohibitedIn\":[],\"associatedTermsOfService\":[]}"), nil
	}

	if uri == "/markets" {
		return []byte("[{\"symbol\": \"4ART-BTC\",\"baseCurrencySymbol\": \"4ART\",\"quoteCurrencySymbol\": \"BTC\",\"minTradeSize\": \"10.00000000\",\"precision\": 8,\"status\": \"ONLINE\",\"createdAt\": \"2020-06-10T15:05:29.833Z\",\"notice\": \"\",\"prohibitedIn\": [\"US\"]},{\"symbol\": \"4ART-USDT\",\"baseCurrencySymbol\": \"4ART\",\"quoteCurrencySymbol\": \"USDT\",\"minTradeSize\": \"10.00000000\",\"precision\": 5,\"status\": \"ONLINE\",\"createdAt\": \"2020-06-10T15:05:40.98Z\",\"notice\": \"\",\"prohibitedIn\": [\"US\"]}]"), nil
	}

	if uri == "/markets/fakesymbol/summary" {
		return []byte("{\"symbol\":\"ETH-BTC\",\"high\":\"0.03894964\",\"low\":\"0.03650000\",\"volume\":\"18494.04035144\",\"quoteVolume\":\"696.42899671\",\"percentChange\":\"-3.33\",\"updatedAt\":\"2020-09-04T04:37:45.107Z\"}"), nil
	}

	if uri == "/markets/summaries" {
		return []byte("[{\"symbol\": \"4ART-BTC\",\"high\": \"0.00000275\",\"low\": \"0.00000249\",\"volume\": \"54499.59344453\",\"quoteVolume\": \"0.13917073\",\"percentChange\": \"10.44\",\"updatedAt\": \"2020-09-04T04:58:55.447Z\"},{\"symbol\": \"4ART-USDT\",\"high\": \"0.02880000\",\"low\": \"0.02667000\",\"volume\": \"48259.53706735\",\"quoteVolume\": \"1320.75839607\",\"percentChange\": \"-6.11\",\"updatedAt\": \"2020-09-04T04:33:20.01Z\"}]"), nil
	}

	if uri == "/markets/fakesymbol/ticker" {
		return []byte("{\"symbol\":\"ETH-BTC\",\"lastTradeRate\":\"0.03760069\",\"bidRate\":\"0.03760103\",\"askRate\":\"0.03762798\"}"), nil
	}

	if uri == "/markets/tickers" {
		return []byte("[{\"symbol\": \"ETH-BTC\",\"lastTradeRate\": \"0.03760069\",\"bidRate\": \"0.03760103\",\"askRate\": \"0.03762798\"},{\"symbol\": \"ETH-FAKE\",\"lastTradeRate\": \"1.03760069\",\"bidRate\": \"1.03760103\",\"askRate\": \"1.03762798\"}]"), nil
	}

	if uri == "/orders/fakeOrder" {
		return []byte("{\"id\": \"55eb2c82-4184-4a24-8b6e-ee154b2f7eaf\",\"marketSymbol\": \"XRP-BTC\",\"direction\": \"BUY\",\"type\": \"LIMIT\",\"quantity\": \"77.53046131\",\"limit\": \"0.00003528\",\"timeInForce\": \"GOOD_TIL_CANCELLED\",\"fillQuantity\": \"77.53046131\",\"commission\": \"0.00000682\",\"proceeds\": \"0.00272829\",\"status\": \"CLOSED\",\"createdAt\": \"2017-10-20T18:27:20.747Z\",\"updatedAt\": \"2017-10-20T18:27:20.763Z\",\"closedAt\": \"2017-10-20T18:27:20.763Z\"}"), nil
	}

	if uri == "/orders/open" {
		return []byte("[{\"id\": \"55eb2c82-4184-4a24-8b6e-ee154b2f7eaf\",\"marketSymbol\": \"XRP-BTC\",\"direction\": \"BUY\",\"type\": \"LIMIT\",\"quantity\": \"77.53046131\",\"limit\": \"0.00003528\",\"timeInForce\": \"GOOD_TIL_CANCELLED\",\"fillQuantity\": \"77.53046131\",\"commission\": \"0.00000682\",\"proceeds\": \"0.00272829\",\"status\": \"OPEN\",\"createdAt\": \"2017-10-20T18:27:20.747Z\",\"updatedAt\": \"2017-10-20T18:27:20.763Z\",\"closedAt\": \"2017-10-20T18:27:20.763Z\"}]"), nil
	}

	if uri == "/orders/closed" {
		return []byte("[{\"id\": \"55eb2c82-4184-4a24-8b6e-ee154b2f7eaf\",\"marketSymbol\": \"XRP-BTC\",\"direction\": \"BUY\",\"type\": \"LIMIT\",\"quantity\": \"77.53046131\",\"limit\": \"0.00003528\",\"timeInForce\": \"GOOD_TIL_CANCELLED\",\"fillQuantity\": \"77.53046131\",\"commission\": \"0.00000682\",\"proceeds\": \"0.00272829\",\"status\": \"CLOSED\",\"createdAt\": \"2017-10-20T18:27:20.747Z\",\"updatedAt\": \"2017-10-20T18:27:20.763Z\",\"closedAt\": \"2017-10-20T18:27:20.763Z\"}]"), nil
	}

	return nil, errors.New("test resource not found")
}

func (this *fakeBittrexClient) authenticate(request *http.Request, payload string, uri string, method string) error {
	return nil
}
