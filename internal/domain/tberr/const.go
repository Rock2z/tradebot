package tberr

import "errors"

var (
	ErrNotFound    = errors.New("not found")
	ErrIndexExceed = errors.New("index exceed")
)
