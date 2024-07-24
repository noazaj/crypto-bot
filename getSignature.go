package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"net/url"
)

func GetKrakenSignature(urlPath string, values url.Values, secret []byte) string {
	sha := sha256.New()
	sha.Write([]byte(values.Get("nonce") + values.Encode()))
	shasum := sha.Sum(nil)

	mac := hmac.New(sha512.New, secret)
	mac.Write(append([]byte(urlPath), shasum...))
	macsum := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(macsum)
}
