package barcode

import (
	"strconv"
	"strings"

	"github.com/neonxp/checksum"
)

// Check barcode
func Check(number string) error {
	sum := 0
	numbers := strings.Split(number, "")
	l := len(numbers)
	for i := l; i > 0; i-- {
		ch := numbers[i-1]
		num, err := strconv.Atoi(string(ch))
		if err != nil {
			return checksum.ErrInvalidNumber
		}
		if (l-i)%2 != 0 {
			sum += num * 3
		} else {
			sum += num
		}
	}
	if sum%10 != 0 {
		return checksum.ErrInvalidChecksum
	}
	return nil
}

func Generate(number string) (string, error) {
	return "", checksum.ErrNotImplemented
}
