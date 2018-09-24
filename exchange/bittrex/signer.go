package bittrex

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"io/ioutil"
	"log"

	ethereum "github.com/ethereum/go-ethereum/common"
)

type Signer struct {
	Key    string `json:"bittrex_key"`
	Secret string `json:"bittrex_secret"`
}

func (s Signer) GetKey() string {
	return s.Key
}

func (s Signer) Sign(msg string) string {
	mac := hmac.New(sha512.New, []byte(s.Secret))
	if _, err := mac.Write([]byte(msg)); err != nil {
		log.Printf("Encode message error: %s", err.Error())
	}
	return ethereum.Bytes2Hex(mac.Sum(nil))
}

func NewSigner(key, secret string) Signer {
	return Signer{key, secret}
}

func NewSignerFromFile(path string) Signer {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	signer := Signer{}
	err = json.Unmarshal(raw, &signer)
	if err != nil {
		panic(err)
	}
	return signer
}
