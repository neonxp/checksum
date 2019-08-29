package isin

import (
	"strconv"
	"strings"

	"github.com/neonxp/checksum"
)

// Check ISIN code
func Check(number string) error {
	sum := 0
	numbersString := ""
	for _, ch := range strings.ToLower(number) {
		if int(ch) >= 'a' && int(ch) <= 'z' {
			numbersString = numbersString + strconv.Itoa(int(ch-'a'+10))
			continue
		}
		numbersString = numbersString + string(ch)
	}
	numbers := strings.Split(numbersString, "")
	l := len(numbers)
	for i := l; i > 0; i-- {
		ch := numbers[i-1]
		num, err := strconv.Atoi(string(ch))
		if err != nil {
			return checksum.ErrInvalidNumber
		}
		if (l-i)%2 != 0 {
			n := num * 2
			if n > 9 {
				sum += n - 9
			} else {
				sum += n
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

func Generate(number string) (string, error) {
	// TODO implement
	return "", checksum.ErrNotImplemented
}
