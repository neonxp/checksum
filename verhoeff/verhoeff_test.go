package verhoeff

import (
	"testing"

	"github.com/neonxp/checksum"
)

func TestVerhoeff(t *testing.T) {
	samples := map[string]error{
		"4561261212345464":        checksum.ErrInvalidChecksum,
		"A561261212345464":        checksum.ErrInvalidNumber,
		"758722":                  nil,
		"123451":                  nil,
		"1428570":                 nil,
		"1234567890120":           nil,
		"84736430954837284567892": nil,
	}
	for num, result := range samples {
		if err := Check(num); err != result {
			t.Errorf("Expected %+v actual %+v", result, err)
		}
	}

	num := "4561261212345467"
	checksum, err := Generate(num)
	if err != nil {
		t.Error(err)
	}
	numberWithChecksum := num + checksum
	if err := Check(numberWithChecksum); err != nil {
		t.Errorf("Expected no error actual %+v", err)
	}
}
