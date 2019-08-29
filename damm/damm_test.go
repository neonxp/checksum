package damm

import (
	"log"
	"testing"

	"github.com/neonxp/checksum"
)

func TestDamm(t *testing.T) {
	samples := map[string]error{
		"456126121234546": checksum.ErrInvalidChecksum,
		"A56126121234546": checksum.ErrInvalidNumber,
		"5724":            nil,
	}
	for num, result := range samples {
		if err := Check(num); err != result {
			t.Errorf("Expected %+v actual %+v for %s", result, err, num)
		}
	}

	num := "572"
	checksum, err := Generate(num)
	if err != nil {
		t.Error(err)
	}
	numberWithChecksum := num + checksum
	log.Println(numberWithChecksum)
	if err := Check(numberWithChecksum); err != nil {
		t.Errorf("Expected no error actual %+v for %s", err, numberWithChecksum)
	}
}
