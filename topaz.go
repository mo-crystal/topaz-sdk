package topazsdk

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Manager struct {
	topazServer string
	privateKey  rsa.PrivateKey
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const VERSION = "v1.0"

func NewManager(serverUrl, _privateKey string) (*Manager, error) {
	resp, err := http.Get(serverUrl + "/")
	if err != nil {
		return nil, errors.New("invalid topaz server")
	}

	_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("invalid topaz server")
	}

	body := Response{}
	err = json.Unmarshal(_body, &body)
	if err != nil {
		return nil, errors.New("invalid topaz server")
	}

	if body.Msg != "topaz server: "+VERSION {
		return nil, errors.New("invalid topaz server")
	}

	key, err := base64.StdEncoding.DecodeString(_privateKey)
	if err != nil {
		return nil, errors.New("invalid private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(key)
	if err != nil {
		return nil, errors.New("invalid private key")
	}

	return &Manager{
		topazServer: serverUrl,
		privateKey:  *privateKey,
	}, nil
}
