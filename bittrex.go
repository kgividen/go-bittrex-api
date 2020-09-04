package main

import (
	"encoding/json"
	"log"

	"github.com/shopspring/decimal"
)

type BittrexAPI struct {
	uri    string
	client Client
}

func NewBittrexAPI(client Client, uri string) *BittrexAPI {
	return &BittrexAPI{client: client, uri: uri}
}

//func (this *BittrexAPI) getResource(resource string, item interface{}) interface{}{
//	uri := this.uri + resource
//	body := this.client.Do("GET", uri, "", false)
//	convertJSON(body, &item)
//	return item
//}

func (this *BittrexAPI) getMarket(symbol string) Market {
	uri := this.uri + "/markets/" + symbol
	body := this.client.Do("GET", uri, "", false)

	market := Market{}
	convertJSON(body, &market)

	return market
}

func (this *BittrexAPI) getMarkets() []Market {
	uri := this.uri + "/markets"
	body := this.client.Do("GET", uri, "", false)

	var markets []Market
	convertJSON(body, &markets)

	return markets
}

func (this *BittrexAPI) getMarketSummary(symbol string) MarketSummary {
	uri := this.uri + "/markets/" + symbol + "/summary"
	body := this.client.Do("GET", uri, "", false)

	marketSummary := MarketSummary{}
	convertJSON(body, &marketSummary)

	return marketSummary
}

func (this *BittrexAPI) getMarketSummaries() []MarketSummary {
	uri := this.uri + "/markets/summaries"
	body := this.client.Do("GET", uri, "", false)

	var marketSummaries []MarketSummary
	convertJSON(body, &marketSummaries)

	return marketSummaries
}

func (this *BittrexAPI) getMarketTicker(symbol string) MarketTicker {
	uri := this.uri + "/markets/" + symbol + "/ticker"
	body := this.client.Do("GET", uri, "", false)

	marketTicker := MarketTicker{}
	convertJSON(body, &marketTicker)

	return marketTicker
}

func (this *BittrexAPI) getMarketTickers() []MarketTicker {
	uri := this.uri + "/markets/tickers"
	body := this.client.Do("GET", uri, "", false)

	var marketTickers []MarketTicker
	convertJSON(body, &marketTickers)

	return marketTickers
}

func (this *BittrexAPI) getCurrency(symbol string) Currency {
	uri := this.uri + "/currencies/" + symbol
	body := this.client.Do("GET", uri, "", false)

	currency := Currency{}
	convertJSON(body, &currency)

	return currency
}

func (this *BittrexAPI) getBalances() []Balance {
	uri := this.uri + "/balances"
	body := this.client.Do("GET", uri, "", true)

	var balances []Balance
	convertJSON(body, &balances)

	return balances
}

func (this *BittrexAPI) getOrder(orderID string) Order {
	uri := this.uri + "/orders/" + orderID
	body := this.client.Do("GET", uri, "", false)

	order := Order{}
	convertJSON(body, &order)

	return order
}

func (this *BittrexAPI) getOrders(openOrClosed string) []Order {
	uri := this.uri + "/orders/" + openOrClosed
	body := this.client.Do("GET", uri, "", true)

	var orders []Order
	convertJSON(body, &orders)

	return orders
}

//func (this *BittrexAPI) getOrdersOpen() {
//	uri := this.uri + "/orders/open"
//	body := this.client.Do("GET", uri, "", true)
//
//	var orders []Order
//	convertJSON(body, &orders)
//
//	return orders
//}
//////////////////////////////////////////
func convertJSON(body []byte, item interface{}) {
	if err := json.Unmarshal(body, &item); err != nil {
		log.Println(err)
	}
}

//////////////////////////////////////////
type Currency struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	CoinType         string `json:"coinType"`
	Status           string `json:"status"`
	MinConfirmations string `json:"minConfirmations"`
	Notice           string `json:"notice"`
	TxFee            string `json:"txFee"`
	LogoUrl          string `json:"logoUrl"`
	ProhibitedIn     string `json:"prohibitedIn"`
	BaseAddress      string `json:"baseAddress"`
}

type Balance struct {
	CurrencySymbol string `json:"currencySymbol"`
	Total          string `json:"total"`
	Available      string `json:"available"`
	UpdatedAt      string `json:"updatedAt"`
}

type Market struct {
	Symbol              string   `json:"symbol"`
	BaseCurrencySymbol  string   `json:"baseCurrencySymbol"`
	QuoteCurrencySymbol string   `json:"quoteCurrencySymbol"`
	MinTradeSize        string   `json:"minTradeSize"`
	Precision           int32    `json:"precision"`
	Status              string   `json:"status"`
	CreatedAt           string   `json:"createdAt"`
	Notice              string   `json:"notice"`
	ProhibitedIn        []string `json:"prohibitedIn"`
}

type MarketSummary struct {
	Symbol        string          `json:"symbol"`
	High          decimal.Decimal `json:"high,string"`
	Low           decimal.Decimal `json:"low,string"`
	Volume        decimal.Decimal `json:"volume,string"`
	QuoteVolume   decimal.Decimal `json:"quoteVolume,string"`
	PercentChange decimal.Decimal `json:"percentChange,string"`
	UpdatedAt     string          `json:"updatedAt"`
}

type MarketTicker struct {
	Symbol        string          `json:"symbol"`
	LastTradeRate decimal.Decimal `json:"lastTradeRate,string"`
	BidRate       decimal.Decimal `json:"bidRate,string"`
	AskRate       decimal.Decimal `json:"askRate,string"`
}

type Order struct {
	OrderID       string          `json:"id"`
	MarketSymbol  string          `json:"marketSymbol"`
	Direction     string          `json:"direction"`
	OrderType     string          `json:"type"`
	Quantity      decimal.Decimal `json:"quantity,string"`
	Limit         decimal.Decimal `json:"limit,string"`
	Ceiling       decimal.Decimal `json:"ceiling,string"`
	TimeInForce   string          `json:"timeInForce"`
	ClientOrderId string          `json:"clientOrderId"`
	FillQuantity  decimal.Decimal `json:"fillQuantity,string"`
	Commission    decimal.Decimal `json:"commission,string"`
	Proceeds      decimal.Decimal `json:"proceeds,string"`
	Status        string          `json:"status"`
	CreatedAt     string          `json:"createdAt"`
	UpdatedAt     string          `json:"updatedAt"`
	ClosedAt      string          `json:"closedAt"`
	OrderToCancel struct {
		OrderType string `json:"type"`
		ID        string `json:"id"`
	} `json:"orderToCancel"`
}
