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
