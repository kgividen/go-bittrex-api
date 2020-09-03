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
	client := &http.Client{}
	api := NewBittrexAPI(client, uri, apiKey, secretKey)
	//log.Println(api.getCurrency("BTC"))
	log.Println(api.getBalances())
}
