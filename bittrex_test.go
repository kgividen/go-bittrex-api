package main

import (
	"io"
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
	bittrex := NewBittrexAPI(&fakeClient{}, "fakeURL", "", "")
	result := bittrex.getCurrency("btc")
	this.So(result, should.Resemble, Currency{})
}

///////////////////////////////////////

type fakeClient struct{}

func (this *fakeClient) Get(url string) (resp *http.Response, err error) {
	return &http.Response{Body: http.NoBody}, nil
}

func (this *fakeClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return &http.Response{}, nil
}
