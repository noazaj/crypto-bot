package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/noazaj/crypto-bot/config"
)

var (
	apiURL = "https://api.kraken.com"
	apiKey = config.ApiKey
	apiSec = config.ApiSec
)

func krakenRequest(uriPath string, data url.Values, key string, secret []byte) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", apiURL, uriPath)

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Error in creating request: %s", err)
	}

	req.Header.Set("API-Key", key)
	req.Header.Set("API-Sign", GetKrakenSignature(uriPath, data, []byte(secret)))

	return req, nil
}

func GetBalance() (string, error) {
	timeData := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	data := url.Values{}
	data.Add("nonce", timeData)
	resp, err := krakenRequest("/0/private/Balance", data, apiKey, []byte(apiSec))
	if err != nil {
		return "", fmt.Errorf("error in getting response: %s", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error in reading body: %s", err)
	}

	return string(body), nil
}

// Kraken Documentation:
// https://docs.kraken.com/rest/#tag/Account-Data/operation/getAccountBalance
// https://github.com/beldur/kraken-go-api-client
