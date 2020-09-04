package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	uri := "https://api.bittrex.com/v3"
	apiKey := os.Getenv("BITTREX_API_KEY")
	secretKey := os.Getenv("BITTREX_SECRET_KEY")
	httpClient := &http.Client{}
	client := newBittrexClient(apiKey, secretKey, httpClient)
	api := NewBittrexAPI(client, uri)
	//log.Println(api.getCurrency("BTC"))
	//log.Println(api.getBalances())
	log.Printf("getMarket: %v", api.getMarket("ETH-BTC"))
}
