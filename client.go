package bittrex

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type bittrexClient struct {
	apiKey    string
	secretKey string
	client    Http
}

func newBittrexClient(apiKey string, secretKey string, client Http) *bittrexClient {
	return &bittrexClient{apiKey: apiKey, secretKey: secretKey, client: client}
}

func (this *bittrexClient) Do(method string, uri string, payload string, authenticate bool) ([]byte, error) {

	request, err := http.NewRequest(method, uri, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	if authenticate {
		if this.authenticate(request, payload, uri, method) != nil {
			return nil, err
		}
	}

	resp, err := this.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func (this *bittrexClient) authenticate(request *http.Request, payload string, uri string, method string) error {
	if len(this.apiKey) == 0 || len(this.secretKey) == 0 {
		err := errors.New("you need to set API Key and API Secret to call this method")
		return err
	}
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	query := request.URL.Query()
	request.URL.RawQuery = query.Encode()

	hash := sha512.New()
	_, err := hash.Write([]byte(payload))
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
	return err
}
