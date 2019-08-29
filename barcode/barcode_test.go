package barcode

import (
	"testing"

	"github.com/neonxp/checksum"
)

func TestBarcode(t *testing.T) {
	samples := map[string]error{
		"4600051000058": checksum.ErrInvalidChecksum,
		"A600051000057": checksum.ErrInvalidNumber,
		"4600051000057": nil,
		"46009333":      nil,
		"041689300494":  nil,
	}
	for num, result := range samples {
		if err := Check(num); err != result {
			t.Errorf("Expected %+v actual %+v for %s", result, err, num)
		}
	}
}
