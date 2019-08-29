# Collection of checksum algorithms on Go

## Luhn algorithm

> The Luhn algorithm or Luhn formula, also known as the "modulus 10" or "mod 10" algorithm, named after its creator, IBM scientist Hans Peter Luhn, is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers, IMEI numbers, National Provider Identifier numbers in the United States, Canadian Social Insurance Numbers, Israel ID Numbers, Greek Social Security Numbers (ΑΜΚΑ), and survey codes appearing on McDonald's, Taco Bell, and Tractor Supply Co. receipts. 

[Wikipedia](https://en.wikipedia.org/wiki/Luhn_algorithm)

Simple implementation on pure Go.

### Usage

```golang
import "github.com/neonxp/checksum/luhn"
...
err := luhn.Check("4561261212345467")
switch err {
    case luhn.ErrInvalidNumber:
    // Not a number
    case luhn.ErrInvalidChecksum:
    // Invalid checksum
    case nil:
    // Valid number
}
```

## Verhoeff algorithm

> The Verhoeff algorithm is a checksum formula for error detection developed by the Dutch mathematician Jacobus Verhoeff and was first published in 1969. It was the first decimal check digit algorithm which detects all single-digit errors, and all transposition errors involving two adjacent digits, which was at the time thought impossible with such a code.

[Wikipedia](https://en.wikipedia.org/wiki/Verhoeff_algorithm)

### Usage

```golang
import "github.com/neonxp/checksum/verhoeff"
...
numberWithoutChecksum := "4561261212345467"
err := verhoeff.Check(number)
switch err {
    case luhn.ErrInvalidNumber:
    // Not a number
    case luhn.ErrInvalidChecksum:
    // Invalid checksum
    case nil:
    // Valid number
}

checksum, err := verhoeff.Generate(numberWithoutChecksum)
if err != nil {
    panic(err)
}
numberWithChecksum := numberWithoutChecksum + checksum
if err := verhoeff.Generate(numberWithChecksum); err != nil {
    panic(err)
}
```
