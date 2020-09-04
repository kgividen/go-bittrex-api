package main

import (
	"io"
	"net/http"
)

type api interface {
	getCurrency(string) (Currency, error)
	getBalances() ([]Balance, error)
}

type Client interface {
	Do(method, uri, payload string, authenticate bool) ([]byte, error)
	authenticate(request *http.Request, payload string, uri string, method string) error
}

type Http interface {
	Get(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
	Do(req *http.Request) (*http.Response, error)
}
