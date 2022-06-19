package tberr

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrIndexExceed      = errors.New("index exceed")
	ErrInvalidParam     = errors.New("invalid param")
	ErrCashNotEnough    = errors.New("cash not enough")
	ErrNoEnoughShare    = errors.New("no enough share")
	ErrInvalidOperation = errors.New("invalid operation")
)
