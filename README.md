# !!!THIS PROJECT IS NOT MAINTAINED ANYMORE!!!

You can use https://github.com/alioygur/is package instead of this. 

# godash

[![Build Status](https://travis-ci.org/alioygur/godash.svg?branch=master)](https://travis-ci.org/alioygur/godash)
[![GoDoc](https://godoc.org/github.com/alioygur/godash?status.svg)](https://godoc.org/github.com/alioygur/godash)
[![Go Report Card](https://goreportcard.com/badge/github.com/alioygur/godash)](https://goreportcard.com/report/github.com/alioygur/godash)

A utility library in Golang inspired by lodash

## installation

`go get gopkg.in/alioygur/godash.v0`

## Contribute

**we are waiting your contribution**

- Report problems
- Add/Suggest new features/recipes
- Improve/fix documentation

Many thanks to our contributors: [contributors](https://github.com/alioygur/godash/graphs/contributors)


## Is* (collection of checking)

An Example;

```go
func IsEmail(str string) bool
```

```go
func ExampleIsEmail() {
	fmt.Println(IsEmail("jhon@example.com"))
	fmt.Println(IsEmail("invalid.com"))
	fmt.Println(IsEmail(`very.(),:;<>[]".VERY."very@\ "very".unusual@strange.example.com`))
	// Output:
	// true
	// false
	// true
}
```

Full list of Is* functions;

```go
func IsASCII(s string) bool
func IsAlpha(s string) bool
func IsAlphanumeric(s string) bool
func IsBase64(s string) bool
func IsByteLength(str string, min, max int) bool
func IsCreditCard(str string) bool
func IsDNSName(str string) bool
func IsDataURI(str string) bool
func IsDialString(str string) bool
func IsDivisibleBy(str, num string) bool
func IsEmail(s string) bool
func IsFilePath(str string) (bool, int)
func IsFloat(str string) bool
func IsFullWidth(str string) bool
func IsHalfWidth(str string) bool
func IsHexadecimal(str string) bool
func IsHexcolor(str string) bool
func IsIP(str string) bool
func IsIPv4(str string) bool
func IsIPv6(str string) bool
func IsISBN(str string, version int) bool
func IsISBN10(str string) bool
func IsISBN13(str string) bool
func IsISO3166Alpha2(str string) bool
func IsISO3166Alpha3(str string) bool
func IsInRange(value, left, right float64) bool
func IsInt(str string) bool
func IsJSON(str string) bool
func IsLatitude(str string) bool
func IsLongitude(str string) bool
func IsLowerCase(str string) bool
func IsMAC(str string) bool
func IsMatches(str, pattern string) bool
func IsMongoID(str string) bool
func IsMultibyte(s string) bool
func IsNatural(value float64) bool
func IsNegative(value float64) bool
func IsNonNegative(value float64) bool
func IsNonPositive(value float64) bool
func IsNull(str string) bool
func IsNumeric(s string) bool
func IsPort(str string) bool
func IsPositive(value float64) bool
func IsPrintableASCII(s string) bool
func IsRGBcolor(str string) bool
func IsRequestURI(rawurl string) bool
func IsRequestURL(rawurl string) bool
func IsSSN(str string) bool
func IsSemver(str string) bool
func IsStringLength(str string, params ...string) bool
func IsStringMatches(s string, params ...string) bool
func IsURL(str string) bool
func IsUTFDigit(s string) bool
func IsUTFLetter(str string) bool
func IsUTFLetterNumeric(s string) bool
func IsUTFNumeric(s string) bool
func IsUUID(str string) bool
func IsUUIDv3(str string) bool
func IsUUIDv4(str string) bool
func IsUUIDv5(str string) bool
func IsUpperCase(str string) bool
func IsVariableWidth(str string) bool
func IsWhole(value float64) bool
```

## To* (collection of converting)

Examples;

```go
func ExampleToBoolean() {
	fmt.Println(ToBoolean("True"))
	fmt.Println(ToBoolean("true"))
	fmt.Println(ToBoolean("1"))
	fmt.Println(ToBoolean("False"))
	fmt.Println(ToBoolean("false"))
	fmt.Println(ToBoolean("0"))
	// Output:
	// true <nil>
	// true <nil>
	// true <nil>
	// false <nil>
	// false <nil>
	// false <nil>
}

func ExampleToCamelCase() {
	fmt.Println(ToCamelCase("camel case"))
	fmt.Println(ToCamelCase("  camel case  "))
	fmt.Println(ToCamelCase("!!!camel case===="))
	fmt.Println(ToCamelCase("camel-case"))
	fmt.Println(ToCamelCase("camel_case"))
	// Output:
	// CamelCase
	// CamelCase
	// CamelCase
	// CamelCase
	// CamelCase
}

func ExampleToSnakeCase() {
	fmt.Println(ToSnakeCase("SnakeCase"))
	fmt.Println(ToSnakeCase("snake case"))
	fmt.Println(ToSnakeCase("  snake case  "))
	fmt.Println(ToSnakeCase("!!!snake case===="))
	fmt.Println(ToSnakeCase("snake-case"))
	fmt.Println(ToSnakeCase("snake_case"))
	// Output:
	// snake_case
	// snake_case
	// snake_case
	// snake_case
	// snake_case
	// snake_case
}
```

Full list of To* functions;

```go
func ToBoolean(str string) (bool, error)
func ToCamelCase(s string) string
func ToFloat(str string) (float64, error)
func ToInt(str string) (int64, error)
func ToJSON(obj interface{}) (string, error)
func ToSnakeCase(str string) string
func ToString(obj interface{}) string
```

for more documentation [godoc](https://godoc.org/github.com/alioygur/godash)

## Thanks & Authors

I use code/got inspiration from these excellent libraries:

- [asaskevich/govalidator](https://github.com/asaskevich/govalidator) [Go] Package of validators and sanitizers for strings, numerics, slices and structs
- [lodash/lodash](https://github.com/lodash/lodash) A modern JavaScript utility library delivering modularity, performance, & extras.
