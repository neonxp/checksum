package isin

import (
	"testing"

	"github.com/neonxp/checksum"
)

func TestISIN(t *testing.T) {
	samples := map[string]error{
		"4600051000058": checksum.ErrInvalidChecksum,
		"?600051000057": checksum.ErrInvalidNumber,
		"A600051000057": checksum.ErrInvalidChecksum,
		"4000000000006": nil,
		"US0378331005":  nil,
	}
	for num, result := range samples {
		if err := Check(num); err != result {
			t.Errorf("Expected %+v actual %+v for %s", result, err, num)
		}
	}
}
