package example_module

import "errors"

var (
	ErrTickerNotFound = errors.New("ticker not found")
)
