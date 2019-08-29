package luhn

import (
	"testing"

	"github.com/neonxp/checksum"
)

func TestLuhn(t *testing.T) {
	samples := map[string]error{
		"4561261212345464": checksum.ErrInvalidChecksum,
		"A561261212345464": checksum.ErrInvalidNumber,
		"4561261212345467": nil,
	}
	for num, result := range samples {
		if err := Check(num); err != result {
			t.Errorf("Expected %+v actual %+v", result, err)
		}
	}
}
