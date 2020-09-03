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

func (this *BittrexAPI) getBalances() string {
	payload := ""
	uri := this.uri + "/balances"
	method := "GET"
	balances := this.client.Do(method, uri, payload, true)

	//currency := Balances{}
	//if err := json.Unmarshal(body, &currency); err != nil {
	//	log.Println(err)
	//}
	//
	log.Println(balances)
	return ""
}
