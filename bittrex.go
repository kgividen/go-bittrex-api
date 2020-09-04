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

func (this *BittrexAPI) getMarket(symbol string) Market {
	uri := this.uri + "/markets/" + symbol
	//uri := this.uri + "/markets"
	body := this.client.Do("GET", uri, "", false)

	market := Market{}
	if err := json.Unmarshal(body, &market); err != nil {
		log.Println(err)
	}

	return market
}

func (this *BittrexAPI) getCurrency(symbol string) Currency {
	uri := this.uri + "/currencies/" + symbol
	body := this.client.Do("GET", uri, "", false)

	currency := Currency{}
	if err := json.Unmarshal(body, &currency); err != nil {
		log.Println(err)
	}

	return currency
}

func (this *BittrexAPI) getBalances() []Balance {
	uri := this.uri + "/balances"
	body := this.client.Do("GET", uri, "", true)

	var balances []Balance
	if err := json.Unmarshal(body, &balances); err != nil {
		log.Println(err)
	}

	return balances
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
