package main

type api interface {
	getCurrency(string) Currency
	getBalances(string)
}

//////////////////////////////////////////
type Currency struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	CoinType         string `json:"coinType"`
	Status           string `json:"status"`
	MinConfirmations string `json:"minConfirmations"`
	Notice           string `json:"notice"`
	TxFee            string `json:"txFee"`
	LogoUrl          string `json:"logoUrl"`
	ProhibitedIn     string `json:"prohibitedIn"`
	BaseAddress      string `json:"baseAddress"`
}
