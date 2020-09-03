package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	url := "https://api.bittrex.com/v3"
	apiKey := os.Getenv("BITTREX_API_KEY")
	secretKey := os.Getenv("BITTREX_SECRET_KEY")
	client := &http.Client{}
	bittrexApi := NewBittrexAPI(client, url, apiKey, secretKey)
	log.Println(bittrexApi.getCurrency("BTC"))
}
