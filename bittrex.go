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

func (this *BittrexAPI) getCurrency(symbol string) Currency {
	body := this.client.Do("GET", this.uri + "/currencies/" + symbol, "", false)

	currency := Currency{}
	if err := json.Unmarshal(body, &currency); err != nil {
		log.Println(err)
	}

	return currency
}

func (this *BittrexAPI) getBalances() []Balance {
	payload := ""
	uri := this.uri + "/balances"
	method := "GET"
	body := this.client.Do(method, uri, payload, true)

	var balances []Balance

	if err := json.Unmarshal(body, &balances); err != nil {
		log.Println(err)
	}

	return balances
}
