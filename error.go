package topazsdk

import "errors"

var (
	ErrWrongPassword      = errors.New("wrong password")
	ErrInvalidTopazServer = errors.New("invalid topaz server")
	ErrUserNotExists      = errors.New("user not exists")
	ErrInvalidParameter   = errors.New("invalid parameter")
	ErrNetworkError       = errors.New("network error")
)
