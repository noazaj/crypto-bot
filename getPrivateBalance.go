package main

import (
	"net/url"

	"github.com/noazaj/crypto-bot/config"
)

const (
	apiURL = "https://api.kraken.com"
	apiKey = config.ApiKey
	apiSec = config.ApiSec
)



func krakenRequest(uriPath string, data url.Values, apiKey, apiSec string) {
	headers := 
}

// Kraken Documentation:
// https://docs.kraken.com/rest/#tag/Account-Data/operation/getAccountBalance
// https://github.com/beldur/kraken-go-api-client
