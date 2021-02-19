package bittrex

import (
	"encoding/json"
	"errors"
)

type BittrexAPI struct {
	uri    string
	client Client
}

func NewBittrexAPI(client Client, uri string) *BittrexAPI {
	return &BittrexAPI{client: client, uri: uri}
}

func (this *BittrexAPI) GetMarket(symbol string) (*Market, error) {
	uri := this.uri + "/markets/" + symbol
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	market := Market{}
	if err := json.Unmarshal(body, &market); err != nil {
		return nil, err
	}

	return &market, nil
}

func (this *BittrexAPI) GetMarkets() ([]*Market, error) {
	uri := this.uri + "/markets"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	var markets []*Market
	if err := json.Unmarshal(body, &markets); err != nil {
		return nil, err
	}

	return markets, nil
}

func (this *BittrexAPI) GetMarketSummary(symbol string) (*MarketSummary, error) {
	uri := this.uri + "/markets/" + symbol + "/summary"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	marketSummary := MarketSummary{}
	if err := json.Unmarshal(body, &marketSummary); err != nil {
		return nil, err
	}

	return &marketSummary, nil
}

func (this *BittrexAPI) GetMarketSummaries() ([]*MarketSummary, error) {
	uri := this.uri + "/markets/summaries"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	var marketSummaries []*MarketSummary
	if err := json.Unmarshal(body, &marketSummaries); err != nil {
		return nil, err
	}

	return marketSummaries, nil
}

func (this *BittrexAPI) GetMarketTicker(symbol string) (*MarketTicker, error) {
	uri := this.uri + "/markets/" + symbol + "/ticker"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	marketTicker := MarketTicker{}
	if err := json.Unmarshal(body, &marketTicker); err != nil {
		return nil, err
	}

	return &marketTicker, nil
}

func (this *BittrexAPI) GetMarketTickers() ([]*MarketTicker, error) {
	uri := this.uri + "/markets/tickers"
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	var marketTickers []*MarketTicker
	if err := json.Unmarshal(body, &marketTickers); err != nil {
		return nil, err
	}

	return marketTickers, nil
}

func (this *BittrexAPI) GetCurrency(symbol string) (*Currency, error) {
	uri := this.uri + "/currencies/" + symbol
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	currency := Currency{}
	if err := json.Unmarshal(body, &currency); err != nil {
		return nil, err
	}

	return &currency, nil
}

func (this *BittrexAPI) GetBalances() ([]*Balance, error) {
	uri := this.uri + "/balances"
	body, err := this.client.Do("GET", uri, "", true)
	if err != nil {
		return nil, err
	}

	var balances []*Balance
	if err := json.Unmarshal(body, &balances); err != nil {
		return nil, err
	}

	return balances, nil
}

func (this *BittrexAPI) GetOrder(orderID string) (*Order, error) {
	uri := this.uri + "/orders/" + orderID
	body, err := this.client.Do("GET", uri, "", false)
	if err != nil {
		return nil, err
	}

	order := Order{}
	if err := json.Unmarshal(body, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (this *BittrexAPI) GetOrders(openOrClosed string) ([]*Order, error) {
	uri := this.uri + "/orders/" + openOrClosed
	body, err := this.client.Do("GET", uri, "", true)
	if err != nil {
		return nil, err
	}

	var orders []*Order
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

//Required marketSymbol, direction, type, timeInForce
func (this *BittrexAPI) CreateOrder(order Order) (*Order, error) {
	payload, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	uri := this.uri + "/orders"
	body, err := this.client.Do("POST", uri, string(payload), true)
	if err != nil {
		return nil, err
	}

	var returnOrder *Order
	if err := json.Unmarshal(body, &returnOrder); err != nil {
		return nil, errors.New(err.Error() + string(body))
	}

	return returnOrder, nil
}
