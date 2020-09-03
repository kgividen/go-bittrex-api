package main

import (
	"log"
	"os"
)

func main() {
	url := "https://api.bittrex.com/v3"
	apiKey := os.Getenv("BITTREX_API_KEY")
	secretKey := os.Getenv("BITTREX_SECRET_KEY")
	bittrexApi := NewBittrexAPI(url, apiKey, secretKey)
	log.Println(bittrexApi.getCurrency("BTC"))
}
