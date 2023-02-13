package topazsdk

import "errors"

var (
	ErrInvalidSelfName    = errors.New("invalid self name(cannot be empty)")
	ErrWrongPassword      = errors.New("wrong password")
	ErrInvalidTopazServer = errors.New("invalid topaz server")
	ErrUserNotExists      = errors.New("user not exists")
	ErrInvalidParameter   = errors.New("invalid parameter")
	ErrNetworkError       = errors.New("network error")
)
