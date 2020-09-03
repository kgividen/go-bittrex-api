package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BittrexAPI struct {
	uri       string
	apiKey    string
	secretKey string
	client    Client
}

func NewBittrexAPI(client Client, uri string, apiKey string, secretKey string) *BittrexAPI {
	return &BittrexAPI{client: client, uri: uri, apiKey: apiKey, secretKey: secretKey}
}

func (this *BittrexAPI) getCurrency(symbol string) Currency {
	resp, err := this.client.Get(this.uri + "/currencies/" + symbol)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

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
	request, err := http.NewRequest(method, uri, strings.NewReader(payload))
	if err != nil {
		return ""
	}

	authNeeded := true
	if authNeeded {
		if len(this.apiKey) == 0 || len(this.secretKey) == 0 {
			err = errors.New("you need to set API Key and API Secret to call this method")
			return ""
		}
		timestamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
		query := request.URL.Query()
		request.URL.RawQuery = query.Encode()

		hash := sha512.New()
		_, err = hash.Write([]byte(payload))
		contentHash := hex.EncodeToString(hash.Sum(nil))

		preSigned := strings.Join([]string{timestamp, uri, method, contentHash}, "")
		sigHash := hmac.New(sha512.New, []byte(this.secretKey))
		_, err = sigHash.Write([]byte(preSigned))
		signature := hex.EncodeToString(sigHash.Sum(nil))

		request.Header.Add("Api-Key", this.apiKey)
		request.Header.Add("Api-Timestamp", timestamp)
		request.Header.Add("Api-Content-Hash", contentHash)
		request.Header.Add("Api-Signature", signature)

		request.Header.Add("Content-Type", "application/json;charset=utf-8")
		request.Header.Add("Accept", "application/json")
	}

	resp, err := this.client.Do(request)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(body)
	//currency := Balances{}
	//if err := json.Unmarshal(body, &currency); err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(currency)
	return ""
}

