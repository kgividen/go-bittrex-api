package main

import (
	"encoding/json"
	"log"
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
	Symbol        string  `json:"symbol"`
	High          float64 `json:"high,string"`
	Low           float64 `json:"low,string"`
	Volume        float64 `json:"volume,string"`
	QuoteVolume   float64 `json:"quoteVolume,string"`
	PercentChange float64 `json:"percentChange,string"`
	UpdatedAt     string  `json:"updatedAt"`
}

type MarketTicker struct {
	Symbol        string  `json:"symbol"`
	LastTradeRate float64 `json:"lastTradeRate,string"`
	BidRate       float64 `json:"bidRate,string"`
	AskRate       float64 `json:"askRate,string"`
}
