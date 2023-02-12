package topazsdk

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string
	Admin    bool   `json:"admin"`
	Email    string `json:"email"`
	Data     string `json:"data"`
	Banned   bool   `json:"banned"`
}

func (m *Manager) PullUser(_uid int, _password ...string) (*User, error) {
	uid := strconv.Itoa(_uid)
	password := ""
	if len(_password) == 1 {
		password = _password[0]
	}

	signature := Sign(m.topazServer+uid+password, m.privateKey)
	if signature == "" {
		return nil, ErrInvalidParameter
	}

	form := url.Values{
		"UserId":    {uid},
		"Server":    {m.topazServer},
		"Password":  {password},
		"Signature": {signature},
	}

	resp, err := http.PostForm(m.topazServer+"/pull-user", form)
	if err != nil {
		return nil, ErrNetworkError
	}

	defer resp.Body.Close()
	_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrNetworkError
	}

	body := Response{}
	err = json.Unmarshal(_body, &body)
	if err != nil {
		return nil, ErrInvalidTopazServer
	}

	if body.Code != 0 {
		return nil, errors.New(body.Msg)
	}

	user, ok := body.Data.(User)
	if !ok {
		return nil, ErrInvalidTopazServer
	}

	if body.Msg == "wrong password" {
		return &user, ErrWrongPassword
	} else {
		return &user, nil
	}
}
