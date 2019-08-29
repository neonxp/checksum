package luhn

import (
	"strconv"

	"github.com/neonxp/checksum"
)

// Check number is correct by luhn algorithm
func Check(number string) error {
	mod := len(number) % 2
	sum := 0
	for i, ch := range number {
		num, err := strconv.Atoi(string(ch))
		if err != nil {
			return checksum.ErrInvalidNumber
		}
		if i%2 == mod {
			if num < 5 {
				sum += num * 2
			} else {
				sum += num*2 - 9
			}
		} else {
			sum += num
		}
	}
	if sum%10 != 0 {
		return checksum.ErrInvalidChecksum
	}
	return nil
}
