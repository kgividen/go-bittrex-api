package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type BittrexAPI struct {
	baseURL   string
	apiKey    string
	secretKey string
	client    Client
}

func NewBittrexAPI(client Client, baseURL string, apiKey string, secretKey string) *BittrexAPI {
	return &BittrexAPI{client: client, baseURL: baseURL, apiKey: apiKey, secretKey: secretKey}
}

func (this *BittrexAPI) getCurrency(symbol string) Currency {
	resp, err := this.client.Get(this.baseURL + "/currencies/" + symbol)
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

func (this *BittrexAPI) getBalances(symbol string) {
	//
	//method := "GET"
	url := "https://api.bittrex.com/v3"
	//payload := ""
	//req, err := http.NewRequest(method, url, strings.NewReader(payload))
	//if err != nil {
	//	return
	//}

	//authNeeded := false
	//if authNeeded {
	//	if len(apiKey) == 0 || len(apiSecret) == 0 {
	//		err = errors.New("you need to set API Key and API Secret to call this method")
	//		return
	//	}
	//	nonce := time.Now().UnixNano()
	//	query := req.URL.Query()
	//	query.Set("apikey", apiKey)
	//	query.Set("nonce", fmt.Sprintf("%d", nonce))
	//	req.URL.RawQuery = query.Encode()
	//	mac := hmac.New(sha512.New, []byte(apiSecret))
	//	_, err = mac.Write([]byte(req.URL.String()))
	//	sig := hex.EncodeToString(mac.Sum(nil))
	//	req.Header.Add("apisign", sig)
	//}

	resp, err := this.client.Get(url + "/currencies/" + symbol)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(body)
	currency := Currency{}
	if err := json.Unmarshal(body, &currency); err != nil {
		log.Println(err)
	}

	log.Println(currency)
}
