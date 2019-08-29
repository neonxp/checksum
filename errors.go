package checksum

import "errors"

var (
	ErrInvalidNumber   = errors.New("invalid number")   // number contains non numeric symbols
	ErrInvalidChecksum = errors.New("invalid checksum") // number not correct by luhn algorithm
)
