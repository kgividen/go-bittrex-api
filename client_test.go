package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/smartystreets/gunit"
)

func TestBittrexClientFixture(t *testing.T) {
	gunit.Run(new(BittrexClientFixture), t)
}

type BittrexClientFixture struct {
	*gunit.Fixture
}

func (this *BittrexClientFixture) Setup() {
}

func (this *BittrexClientFixture) Test() {
}

type fakeHttpClient struct{}

func (this *fakeHttpClient) Get(url string) (resp *http.Response, err error) {
	return &http.Response{Body: http.NoBody}, nil
}

func (this *fakeHttpClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return &http.Response{}, nil
}
func (this *fakeHttpClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}
