package topazsdk

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Manager struct {
	topazServer string
	privateKey  string
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const VERSION = "v1.0"

func NewManager(serverUrl, privateKey string) (*Manager, error) {
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

	return &Manager{
		topazServer: serverUrl,
		privateKey:  privateKey,
	}, nil
}
