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
		return Market{}, err
	}

	market := Market{}
	if err := json.Unmarshal(body, &market); err != nil {
		return Market{}, err
	}

	return market, nil
}

func (this *BittrexAPI) getMarkets() ([]Market, error) {
	uri := this.uri + "/markets"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	var markets []Market
	if err := json.Unmarshal(body, &markets); err != nil {
		return nil, err
	}

	return markets, nil
}

func (this *BittrexAPI) getMarketSummary(symbol string) (MarketSummary, error) {
	uri := this.uri + "/markets/" + symbol + "/summary"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return MarketSummary{}, err
	}

	marketSummary := MarketSummary{}
	if err := json.Unmarshal(body, &marketSummary); err != nil {
		return MarketSummary{}, err
	}

	return marketSummary, nil
}

func (this *BittrexAPI) getMarketSummaries() ([]MarketSummary, error) {
	uri := this.uri + "/markets/summaries"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	var marketSummaries []MarketSummary
	if err := json.Unmarshal(body, &marketSummaries); err != nil {
		return nil, err
	}

	return marketSummaries, nil
}

func (this *BittrexAPI) getMarketTicker(symbol string) (MarketTicker, error) {
	uri := this.uri + "/markets/" + symbol + "/ticker"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return MarketTicker{}, err
	}

	marketTicker := MarketTicker{}
	if err := json.Unmarshal(body, &marketTicker); err != nil {
		return MarketTicker{}, err
	}

	return marketTicker, nil
}

func (this *BittrexAPI) getMarketTickers() ([]MarketTicker, error) {
	uri := this.uri + "/markets/tickers"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	var marketTickers []MarketTicker
	if err := json.Unmarshal(body, &marketTickers); err != nil {
		return nil, err
	}

	return marketTickers, nil
}

func (this *BittrexAPI) getCurrency(symbol string) (Currency, error) {
	uri := this.uri + "/currencies/" + symbol
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return Currency{}, err
	}

	currency := Currency{}
	if err := json.Unmarshal(body, &currency); err != nil {
		return Currency{}, err
	}

	return currency, nil
}

func (this *BittrexAPI) getBalances() ([]Balance, error) {
	uri := this.uri + "/balances"
	body, err := this.client.Do("GET", uri, "", true)
	if err != nil {
		return nil, err
	}

	var balances []Balance
	if err := json.Unmarshal(body, &balances); err != nil {
		return nil, err
	}

	return balances, nil
}

func (this *BittrexAPI) getOrder(orderID string) (Order, error) {
	uri := this.uri + "/orders/" + orderID
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return Order{}, err
	}

	order := Order{}
	if err := json.Unmarshal(body, &order); err != nil {
		return Order{}, err
	}

	return order, nil
}

func (this *BittrexAPI) getOrders(openOrClosed string) ([]Order, error) {
	uri := this.uri + "/orders/" + openOrClosed
	body, err := this.client.Do("GET", uri, "", true)
	if err != nil {
		return nil, err
	}

	var orders []Order
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

//Required marketSymbol, direction, type, timeInForce
func (this *BittrexAPI) createOrder(order Order) ([]Order, error) {
	payload, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	uri := this.uri + "/orders/"
	body, err := this.client.Do("POST", uri, string(payload), true)
	if err != nil {
		return nil, err
	}

	var orders []Order
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, err
	}

	return orders, nil
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
	Symbol        string `json:"symbol"`
	High          *Dec   `json:"high,string"`
	Low           *Dec   `json:"low,string"`
	Volume        *Dec   `json:"volume,string"`
	QuoteVolume   *Dec   `json:"quoteVolume,string"`
	PercentChange *Dec   `json:"percentChange,string"`
	UpdatedAt     string `json:"updatedAt"`
}

type MarketTicker struct {
	Symbol        string `json:"symbol"`
	LastTradeRate *Dec   `json:"lastTradeRate,string"`
	BidRate       *Dec   `json:"bidRate,string"`
	AskRate       *Dec   `json:"askRate,string"`
}

type Order struct {
	OrderID       string       `json:"id,omitempty"`
	MarketSymbol  string       `json:"marketSymbol"` //Required
	Direction     string       `json:"direction"`    //Required - Buy, Sell
	OrderType     string       `json:"type"`         //Required - LIMIT, MARKET, CEILING_LIMIT, CEILING_MARKET
	Quantity      *Dec         `json:"quantity,string,omitempty"`
	Limit         *Dec         `json:"limit,string,omitempty"`
	Ceiling       *Dec         `json:"ceiling,string,omitempty"`
	TimeInForce   string       `json:"timeInForce,omitempty"` //GOOD_TIL_CANCELLED, IMMEDIATE_OR_CANCEL, FILL_OR_KILL, POST_ONLY_GOOD_TIL_CANCELLED, BUY_NOW
	ClientOrderId string       `json:"clientOrderId,omitempty"`
	FillQuantity  *Dec         `json:"fillQuantity,string,omitempty"`
	Commission    *Dec         `json:"commission,string,omitempty"`
	Proceeds      *Dec         `json:"proceeds,string,omitempty"`
	Status        string       `json:"status,omitempty"`
	CreatedAt     string       `json:"createdAt,omitempty"`
	UpdatedAt     string       `json:"updatedAt,omitempty"`
	ClosedAt      string       `json:"closedAt,omitempty"`
	UseAwards     bool         `json:"useAwards,omitempty"`
	OrderToCancel *OrderCancel `json:"orderToCancel,omitempty"` //Required -  GOOD_TIL_CANCELLED, IMMEDIATE_OR_CANCEL, FILL_OR_KILL, POST_ONLY_GOOD_TIL_CANCELLED, BUY_NOW
}

type OrderCancel struct {
	OrderType string `json:"type,omitempty"`
	ID        string `json:"id,omitempty"`
}

type Dec struct {
	decimal.Decimal
}
