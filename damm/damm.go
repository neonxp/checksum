package damm

import (
	"strconv"

	"github.com/neonxp/checksum"
)

var d = [10][10]int{
	{0, 3, 1, 7, 5, 9, 8, 6, 4, 2},
	{7, 0, 9, 2, 1, 5, 4, 8, 6, 3},
	{4, 2, 0, 6, 8, 7, 1, 3, 5, 9},
	{1, 7, 5, 0, 9, 8, 3, 4, 2, 6},
	{6, 1, 2, 3, 0, 4, 5, 9, 7, 8},
	{3, 6, 7, 4, 2, 0, 9, 5, 8, 1},
	{5, 8, 6, 9, 7, 2, 0, 1, 3, 4},
	{8, 9, 4, 5, 3, 6, 2, 0, 1, 7},
	{9, 4, 3, 8, 6, 1, 7, 2, 0, 5},
	{2, 5, 8, 1, 4, 3, 6, 7, 9, 0},
}

// Check number is correct by damm algorithm
func Check(number string) error {
	c := 0
	for _, ch := range number {
		num, err := strconv.Atoi(string(ch))
		if err != nil {
			return checksum.ErrInvalidNumber
		}
		c = d[c][num]
	}
	if c != 0 {
		return checksum.ErrInvalidChecksum
	}
	return nil
}

// Generate checksum (must be added to number at right)
func Generate(number string) (string, error) {
	c := 0
	for _, ch := range number {
		num, err := strconv.Atoi(string(ch))
		if err != nil {
			return "", checksum.ErrInvalidNumber
		}
		c = d[c][num]
	}
	return strconv.Itoa(c), nil
}
