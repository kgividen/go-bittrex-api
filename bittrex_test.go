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

///////////////////////////////////////

type fakeBittrexClient struct{}

func (this *fakeBittrexClient) Do(method, uri, payload string, authenticate bool) []byte {
	return nil
}

func (this *fakeBittrexClient) authenticate(request *http.Request, payload string, uri string, method string) error {
	return nil
}
