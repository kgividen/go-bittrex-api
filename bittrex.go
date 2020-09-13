package bittrex

import (
	"encoding/json"
	"errors"

	crypto "github.com/kgividen/crypto_contracts"
)

type BittrexAPI struct {
	uri    string
	client crypto.Client
}

func NewBittrexAPI(client crypto.Client, uri string) *BittrexAPI {
	return &BittrexAPI{client: client, uri: uri}
}

func (this *BittrexAPI) GetMarket(symbol string) (*crypto.Market, error) {
	uri := this.uri + "/markets/" + symbol
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	market := crypto.Market{}
	if err := json.Unmarshal(body, &market); err != nil {
		return nil, err
	}

	return &market, nil
}

func (this *BittrexAPI) GetMarkets() ([]*crypto.Market, error) {
	uri := this.uri + "/markets"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	var markets []*crypto.Market
	if err := json.Unmarshal(body, &markets); err != nil {
		return nil, err
	}

	return markets, nil
}

func (this *BittrexAPI) GetMarketSummary(symbol string) (*crypto.MarketSummary, error) {
	uri := this.uri + "/markets/" + symbol + "/summary"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	marketSummary := crypto.MarketSummary{}
	if err := json.Unmarshal(body, &marketSummary); err != nil {
		return nil, err
	}

	return &marketSummary, nil
}

func (this *BittrexAPI) GetMarketSummaries() ([]*crypto.MarketSummary, error) {
	uri := this.uri + "/markets/summaries"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	var marketSummaries []*crypto.MarketSummary
	if err := json.Unmarshal(body, &marketSummaries); err != nil {
		return nil, err
	}

	return marketSummaries, nil
}

func (this *BittrexAPI) GetMarketTicker(symbol string) (*crypto.MarketTicker, error) {
	uri := this.uri + "/markets/" + symbol + "/ticker"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	marketTicker := crypto.MarketTicker{}
	if err := json.Unmarshal(body, &marketTicker); err != nil {
		return nil, err
	}

	return &marketTicker, nil
}

func (this *BittrexAPI) GetMarketTickers() ([]*crypto.MarketTicker, error) {
	uri := this.uri + "/markets/tickers"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	var marketTickers []*crypto.MarketTicker
	if err := json.Unmarshal(body, &marketTickers); err != nil {
		return nil, err
	}

	return marketTickers, nil
}

func (this *BittrexAPI) GetCurrency(symbol string) (*crypto.Currency, error) {
	uri := this.uri + "/currencies/" + symbol
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	currency := crypto.Currency{}
	if err := json.Unmarshal(body, &currency); err != nil {
		return nil, err
	}

	return &currency, nil
}

func (this *BittrexAPI) GetBalances() ([]*crypto.Balance, error) {
	uri := this.uri + "/balances"
	body, err := this.client.Do("GET", uri, "", true)
	if err != nil {
		return nil, err
	}

	var balances []*crypto.Balance
	if err := json.Unmarshal(body, &balances); err != nil {
		return nil, err
	}

	return balances, nil
}

func (this *BittrexAPI) GetOrder(orderID string) (*crypto.Order, error) {
	uri := this.uri + "/orders/" + orderID
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	order := crypto.Order{}
	if err := json.Unmarshal(body, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (this *BittrexAPI) GetOrders(openOrClosed string) ([]*crypto.Order, error) {
	uri := this.uri + "/orders/" + openOrClosed
	body, err := this.client.Do("GET", uri, "", true)
	if err != nil {
		return nil, err
	}

	var orders []*crypto.Order
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

//Required marketSymbol, direction, type, timeInForce
func (this *BittrexAPI) CreateOrder(order crypto.Order) (*crypto.Order, error) {
	payload, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	uri := this.uri + "/orders"
	body, err := this.client.Do("POST", uri, string(payload), true)
	if err != nil {
		return nil, err
	}

	var returnOrder *crypto.Order
	if err := json.Unmarshal(body, &returnOrder); err != nil {
		return nil, errors.New(err.Error() + string(body))
	}

	return returnOrder, nil
}
