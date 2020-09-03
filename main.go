package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	getCurrency("btc")
}

func getCurrency(symbol string) {
	//apiKey := os.Getenv("BITTREX-API")
	//apiSecret := os.Getenv("BITTREX-SECRET")
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

	resp, err := http.Get(url + "/currencies/" + symbol)
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
