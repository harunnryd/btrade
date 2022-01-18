package core

import "errors"

var (
	ErrDBConn         = errors.New("104002")
	ErrUnauthorized   = errors.New("104003")
	ErrInvalidHeader  = errors.New("104004")
	ErrInvalidPayload = errors.New("104005")
)
