package bittrex

import (
	"io"
	"net/http"

	"github.com/shopspring/decimal"
)

type CryptoAPI interface {
	GetCurrency(string) (*Currency, error)
	GetBalances() ([]*Balance, error)
	GetMarketSummaries() ([]*MarketSummary, error)
	CreateOrder(order Order) (*Order, error)
}

type Client interface {
	Do(method, uri, payload string, authenticate bool) ([]byte, error)
	Authenticate(request *http.Request, payload string, uri string, method string) error
}

type Http interface {
	Get(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
	Do(req *http.Request) (*http.Response, error)
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

type CurrencyPurchase struct {
	Symbol           string `json:"symbol"`
	Price            *Dec   `json:"price"`
	PurchaseDate     string `json:"purchase_date"`
}

type Balance struct {
	CurrencySymbol string `json:"currencySymbol"`
	Total          string `json:"total"`
	Available      string `json:"available"`
	UpdatedAt      string `json:"updatedAt"`
}

type Market struct {
	Symbol              string   `json:"symbol"`
	BaseCurrencySymbol  string   `json:"baseCurrencySymbol"`
	QuoteCurrencySymbol string   `json:"quoteCurrencySymbol"`
	MinTradeSize        string   `json:"minTradeSize"`
	Precision           int32    `json:"precision"`
	Status              string   `json:"status"`
	CreatedAt           string   `json:"createdAt"`
	Notice              string   `json:"notice"`
	ProhibitedIn        []string `json:"prohibitedIn"`
}

type MarketSummary struct {
	Symbol        string `json:"symbol"`
	High          *Dec   `json:"high,string"`
	Low           *Dec   `json:"low,string"`
	Volume        *Dec   `json:"volume,string"`
	QuoteVolume   *Dec   `json:"quoteVolume,string"`
	PercentChange *Dec   `json:"percentChange,string"`
	UpdatedAt     string `json:"updatedAt"`
}

type MarketTicker struct {
	Symbol        string `json:"symbol"`
	LastTradeRate *Dec   `json:"lastTradeRate,string"`
	BidRate       *Dec   `json:"bidRate,string"`
	AskRate       *Dec   `json:"askRate,string"`
}

type Order struct {
	OrderID       string       `json:"id,omitempty"`
	MarketSymbol  string       `json:"marketSymbol"` //Required
	Direction     string       `json:"direction"`    //Required - Buy, Sell
	OrderType     string       `json:"type"`         //Required - LIMIT, MARKET, CEILING_LIMIT, CEILING_MARKET
	Quantity      *Dec         `json:"quantity,string,omitempty"`
	Limit         *Dec         `json:"limit,string,omitempty"`
	Ceiling       *Dec         `json:"ceiling,string,omitempty"`
	TimeInForce   string       `json:"timeInForce,omitempty"` //GOOD_TIL_CANCELLED, IMMEDIATE_OR_CANCEL, FILL_OR_KILL, POST_ONLY_GOOD_TIL_CANCELLED, BUY_NOW
	ClientOrderId string       `json:"clientOrderId,omitempty"`
	FillQuantity  *Dec         `json:"fillQuantity,string,omitempty"`
	Commission    *Dec         `json:"commission,string,omitempty"`
	Proceeds      *Dec         `json:"proceeds,string,omitempty"`
	Status        string       `json:"status,omitempty"`
	CreatedAt     string       `json:"createdAt,omitempty"`
	UpdatedAt     string       `json:"updatedAt,omitempty"`
	ClosedAt      string       `json:"closedAt,omitempty"`
	UseAwards     bool         `json:"useAwards,omitempty"`
	OrderToCancel *OrderCancel `json:"orderToCancel,omitempty"` //Required -  GOOD_TIL_CANCELLED, IMMEDIATE_OR_CANCEL, FILL_OR_KILL, POST_ONLY_GOOD_TIL_CANCELLED, BUY_NOW
}

type OrderCancel struct {
	OrderType string `json:"type,omitempty"`
	ID        string `json:"id,omitempty"`
}

type Dec struct {
	decimal.Decimal
}
