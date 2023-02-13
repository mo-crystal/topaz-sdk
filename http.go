package topazsdk

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func GetToStruct(url string, dst interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return ErrNetworkError
	}

	defer resp.Body.Close()
	_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ErrNetworkError
	}

	err = json.Unmarshal(_body, dst)
	if err != nil {
		return ErrInvalidTopazServer
	} else {
		return nil
	}
}

func PostToStruct(url string, formData url.Values, dst interface{}) error {
	resp, err := http.PostForm(url, formData)
	if err != nil {
		return ErrNetworkError
	}

	defer resp.Body.Close()
	_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ErrNetworkError
	}

	err = json.Unmarshal(_body, dst)
	if err != nil {
		return ErrInvalidTopazServer
	} else {
		return nil
	}
}
