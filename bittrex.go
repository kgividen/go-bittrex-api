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

func (this *BittrexAPI) getMarket(symbol string) (Market, error) {
	uri := this.uri + "/markets/" + symbol
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		log.Println(err)
	}

	market := Market{}
	convertJSON(body, &market)

	return market, nil
}

func (this *BittrexAPI) getMarkets() ([]Market, error) {
	uri := this.uri + "/markets"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		log.Println(err)
	}

	var markets []Market
	convertJSON(body, &markets)

	return markets, nil
}

func (this *BittrexAPI) getMarketSummary(symbol string) (MarketSummary, error) {
	uri := this.uri + "/markets/" + symbol + "/summary"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		log.Println(err)
	}

	marketSummary := MarketSummary{}
	convertJSON(body, &marketSummary)

	return marketSummary, nil
}

func (this *BittrexAPI) getMarketSummaries() ([]MarketSummary, error) {
	uri := this.uri + "/markets/summaries"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		log.Println(err)
	}

	var marketSummaries []MarketSummary
	convertJSON(body, &marketSummaries)

	return marketSummaries, nil
}

func (this *BittrexAPI) getMarketTicker(symbol string) (MarketTicker, error) {
	uri := this.uri + "/markets/" + symbol + "/ticker"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		log.Println(err)
	}

	marketTicker := MarketTicker{}
	convertJSON(body, &marketTicker)

	return marketTicker, nil
}

func (this *BittrexAPI) getMarketTickers() ([]MarketTicker, error) {
	uri := this.uri + "/markets/tickers"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		log.Println(err)
	}

	var marketTickers []MarketTicker
	convertJSON(body, &marketTickers)

	return marketTickers, nil
}

func (this *BittrexAPI) getCurrency(symbol string) (Currency, error) {
	uri := this.uri + "/currencies/" + symbol
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		log.Println(err)
	}

	currency := Currency{}
	convertJSON(body, &currency)

	return currency, nil
}

func (this *BittrexAPI) getBalances() ([]Balance, error) {
	uri := this.uri + "/balances"
	body, err := this.client.Do("GET", uri, "", true)
	if err != nil {
		log.Println(err)
	}

	var balances []Balance
	convertJSON(body, &balances)

	return balances, nil
}

func (this *BittrexAPI) getOrder(orderID string) (Order, error) {
	uri := this.uri + "/orders/" + orderID
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		log.Println(err)
	}

	order := Order{}
	convertJSON(body, &order)

	return order, nil
}

func (this *BittrexAPI) getOrders(openOrClosed string) ([]Order, error) {
	uri := this.uri + "/orders/" + openOrClosed
	body, err := this.client.Do("GET", uri, "", true)
	if err != nil {
		log.Println(err)
	}

	var orders []Order
	convertJSON(body, &orders)

	return orders, nil
}

//OrderType:LIMIT, MARKET, CEILING_LIMIT, CEILING_MARKET
func (this *BittrexAPI) createOrder(order Order) []Order {
	payload, err := json.Marshal(order)

	if err != nil {
		log.Printf("Error creating order json: %v", order)
	}
	uri := this.uri + "/orders/"
	body, err := this.client.Do("POST", uri, string(payload), true)
	if err != nil {
		log.Println(err)
		return nil
	}

	var orders []Order
	convertJSON(body, &orders)

	return orders
}

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
	UseAwards     bool            `json:"useAwards"`
	OrderToCancel struct {
		OrderType string `json:"type"`
		ID        string `json:"id"`
	} `json:"orderToCancel"`
}
