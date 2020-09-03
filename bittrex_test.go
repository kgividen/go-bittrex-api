package main

import (
	"net/http"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestBittrexAPIFixture(t *testing.T) {
	gunit.Run(new(BittrexAPIFixture), t)

}

type BittrexAPIFixture struct {
	*gunit.Fixture
}

func (this *BittrexAPIFixture) Setup() {}

func (this *BittrexAPIFixture) TestGetCurrency() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "fakeURL")
	result := bittrex.getCurrency("btc")
	this.So(result, should.Resemble, Currency{})
}

func (this *BittrexAPIFixture) TestGetBalances() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result := bittrex.getBalances()
	this.So(result, should.Resemble, []Balance{
		{
			CurrencySymbol: "BTC",
			Total:          "0.00000000",
			Available:      "0.00000000",
			UpdatedAt:      "2019-10-29T20:25:10.16Z",
		},{
			CurrencySymbol: "LTC",
			Total:          "0",
			Available:      "0",
			UpdatedAt:      "2020-09-03T21:27:53.8210894Z",
		},
	})
}

///////////////////////////////////////

type fakeBittrexClient struct{}

func (this *fakeBittrexClient) Do(method, uri, payload string, authenticate bool) []byte {
	if uri == "/balances" {
		return []byte("[{\"currencySymbol\": \"BTC\",\"total\": \"0.00000000\",\"available\": \"0.00000000\",\"updatedAt\": \"2019-10-29T20:25:10.16Z\"},{\"currencySymbol\": \"LTC\",\"total\": \"0\",\"available\": \"0\",\"updatedAt\": \"2020-09-03T21:27:53.8210894Z\"}]")
	}
	return nil
}

func (this *fakeBittrexClient) authenticate(request *http.Request, payload string, uri string, method string) error {
	return nil
}
