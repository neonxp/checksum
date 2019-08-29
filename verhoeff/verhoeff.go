package verhoeff

import (
	"strconv"
	"strings"

	"github.com/neonxp/checksum"
)

var d = [10][10]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	{1, 2, 3, 4, 0, 6, 7, 8, 9, 5},
	{2, 3, 4, 0, 1, 7, 8, 9, 5, 6},
	{3, 4, 0, 1, 2, 8, 9, 5, 6, 7},
	{4, 0, 1, 2, 3, 9, 5, 6, 7, 8},
	{5, 9, 8, 7, 6, 0, 4, 3, 2, 1},
	{6, 5, 9, 8, 7, 1, 0, 4, 3, 2},
	{7, 6, 5, 9, 8, 2, 1, 0, 4, 3},
	{8, 7, 6, 5, 9, 3, 2, 1, 0, 4},
	{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
}

var p = [8][10]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	{1, 5, 7, 6, 2, 8, 3, 0, 9, 4},
	{5, 8, 0, 3, 7, 9, 6, 1, 4, 2},
	{8, 9, 1, 6, 0, 4, 3, 5, 2, 7},
	{9, 4, 5, 3, 1, 2, 6, 8, 7, 0},
	{4, 2, 8, 6, 5, 7, 3, 9, 0, 1},
	{2, 7, 9, 3, 8, 0, 6, 4, 1, 5},
	{7, 0, 4, 6, 9, 1, 3, 2, 5, 8},
}

var inv = [10]int{0, 4, 3, 2, 1, 5, 6, 7, 8, 9}

// Check number is correct by luhn algorithm
func Check(number string) error {
	c := 0
	numbers := strings.Split(number, "")
	l := len(numbers)
	for i := l; i > 0; i-- {
		ch := numbers[i-1]
		num, err := strconv.Atoi(string(ch))
		if err != nil {
			return checksum.ErrInvalidNumber
		}
		c = d[c][p[(l-i)&7][num]]
	}
	if c != 0 {
		return checksum.ErrInvalidChecksum
	}
	return nil
}

// Generate checksum (must be added to number at right)
func Generate(number string) (string, error) {
	c := 0
	numbers := strings.Split(number, "")
	l := len(numbers)
	for i := l; i > 0; i-- {
		ch := numbers[i-1]
		num, err := strconv.Atoi(string(ch))
		if err != nil {
			return "", checksum.ErrInvalidNumber
		}
		c = d[c][p[(l-i+1)&7][num]]
	}
	return strconv.Itoa(inv[c]), nil
}
