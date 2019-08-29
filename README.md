# Collection of checksum algorithms on Go

Pure Go implementations.

## ISIN 

> An International Securities Identification Number (ISIN) uniquely identifies a security. Its structure is defined in ISO 6166. The ISIN code is a 12-character alphanumeric code that serves for uniform identification of a security through normalization of the assigned National Number, where one exists, at trading and settlement.

[Wikipedia](https://en.wikipedia.org/wiki/International_Securities_Identification_Number)

### Usage

```golang
import (
    "github.com/neonxp/checksum"
    "github.com/neonxp/checksum/isin"
)
...
err := isin.Check("4000000000006") // 
switch err {
    case checksum.ErrInvalidNumber:
    // Not a number
    case checksum.ErrInvalidChecksum:
    // Invalid checksum
    case nil:
    // Valid number
}
```

## Barcode EAN-8, UPC-12, EAN-13

Validate barcode's checksum

### Usage

```golang
import (
    "github.com/neonxp/checksum"
    "github.com/neonxp/checksum/barcode"
)
...
err := barcode.Check("041689300494") // UPC-12 barcode
switch err {
    case checksum.ErrInvalidNumber:
    // Not a number
    case checksum.ErrInvalidChecksum:
    // Invalid checksum
    case nil:
    // Valid number
}
```

## Luhn algorithm

> The Luhn algorithm or Luhn formula, also known as the "modulus 10" or "mod 10" algorithm, named after its creator, IBM scientist Hans Peter Luhn, is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers, IMEI numbers, National Provider Identifier numbers in the United States, Canadian Social Insurance Numbers, Israel ID Numbers, Greek Social Security Numbers (ΑΜΚΑ), and survey codes appearing on McDonald's, Taco Bell, and Tractor Supply Co. receipts. 

[Wikipedia](https://en.wikipedia.org/wiki/Luhn_algorithm)

### Usage

```golang
import (
    "github.com/neonxp/checksum"
    "github.com/neonxp/checksum/luhn"
)
...
err := luhn.Check("4561261212345467")
switch err {
    case checksum.ErrInvalidNumber:
    // Not a number
    case checksum.ErrInvalidChecksum:
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
import (
    "github.com/neonxp/checksum"
    "github.com/neonxp/checksum/verhoeff"
)
...
numberWithoutChecksum := "4561261212345467"
err := verhoeff.Check(number)
switch err {
    case checksum.ErrInvalidNumber:
    // Not a number
    case checksum.ErrInvalidChecksum:
    // Invalid checksum
    case nil:
    // Valid number
}

checksum, err := verhoeff.Generate(numberWithoutChecksum)
if err != nil {
    panic(err)
}
numberWithChecksum := numberWithoutChecksum + checksum
if err := verhoeff.Check(numberWithChecksum); err != nil {
    panic(err)
}
```

# Damm algorithm

> In error detection, the Damm algorithm is a check digit algorithm that detects all single-digit errors and all adjacent transposition errors. It was presented by H. Michael Damm in 2004.

[Wikipedia](https://en.wikipedia.org/wiki/Damm_algorithm)

### Usage

```golang
import (
    "github.com/neonxp/checksum"
    "github.com/neonxp/checksum/damm"
)
...
numberWithoutChecksum := "4561261212345467"
err := damm.Check(number)
switch err {
    case checksum.ErrInvalidNumber:
    // Not a number
    case checksum.ErrInvalidChecksum:
    // Invalid checksum
    case nil:
    // Valid number
}

checksum, err := damm.Generate(numberWithoutChecksum)
if err != nil {
    panic(err)
}
numberWithChecksum := numberWithoutChecksum + checksum
if err := damm.Check(numberWithChecksum); err != nil {
    panic(err)
}
```
