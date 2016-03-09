# godash
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
var validEmail = "jhon@example.com"
var invalidEmail = "blabla"

godash.IsEmail(validEmail) // result: true
godash.IsEmail(invalidEmail) // result: false
```

Full list of Is* functions;

```go
func IsASCII(str string) bool
func IsAlpha(str string) bool
func IsAlphanumeric(str string) bool
func IsBase64(str string) bool
func IsByteLength(str string, min, max int) bool
func IsCreditCard(str string) bool
func IsDNSName(str string) bool
func IsDataURI(str string) bool
func IsDialString(str string) bool
func IsDivisibleBy(str, num string) bool
func IsEmail(str string) bool
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
func IsMultibyte(str string) bool
func IsNatural(value float64) bool
func IsNegative(value float64) bool
func IsNonNegative(value float64) bool
func IsNonPositive(value float64) bool
func IsNull(str string) bool
func IsNumeric(str string) bool
func IsPort(str string) bool
func IsPositive(value float64) bool
func IsPrintableASCII(str string) bool
func IsRGBcolor(str string) bool
func IsRequestURI(rawurl string) bool
func IsRequestURL(rawurl string) bool
func IsSSN(str string) bool
func IsSemver(str string) bool
func IsStringLength(str string, params ...string) bool
func IsStringMatches(s string, params ...string) bool
func IsURL(str string) bool
func IsUTFDigit(str string) bool
func IsUTFLetter(str string) bool
func IsUTFLetterNumeric(str string) bool
func IsUTFNumeric(str string) bool
func IsUUID(str string) bool
func IsUUIDv3(str string) bool
func IsUUIDv4(str string) bool
func IsUUIDv5(str string) bool
func IsUpperCase(str string) bool
func IsVariableWidth(str string) bool
func IsWhole(value float64) bool
```

## To* (collection of converting)

An Example;

```go
godash.ToBoolean("True") // result: true
godash.ToBoolean("true") // result: true
godash.ToBoolean("1") // result: true
godash.ToBoolean("false") // result: false
godash.ToBoolean("0") // result: false
godash.ToBoolean("blabla") // result: false
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

## Utils (collection of utilities)

```go
func Abs(value float64) float64
func BlackList(str, chars string) string
func ByteLength(str string, params ...string) bool
func Contains(str, substring string) bool
func Count(array []interface{}, iterator ConditionIterator) int
func Each(array []interface{}, iterator Iterator)
func Filter(array []interface{}, iterator ConditionIterator) []interface{}
func Find(array []interface{}, iterator ConditionIterator) interface{}
func GetLine(s string, index int) (string, error)
func GetLines(s string) []string
func LeftTrim(str, chars string) string
func Map(array []interface{}, iterator ResultIterator) []interface{}
func NormalizeEmail(str string) (string, error)
func RemoveTags(s string) string
func ReplacePattern(str, pattern, replace string) string
func Reverse(s string) string
func RightTrim(str, chars string) string
func SafeFileName(str string) string
func Sign(value float64) float64
func StripLow(str string, keepNewLines bool) string
func Trim(str, chars string) string
func Truncate(str string, length int, ending string) string
func WhiteList(str, chars string) string
```

for more documentation [godoc](https://godoc.org/github.com/alioygur/gores)

## Thanks & Authors

I use code/got inspiration from these excellent libraries:

- [asaskevich/govalidator](https://github.com/asaskevich/govalidator) [Go] Package of validators and sanitizers for strings, numerics, slices and structs
- [lodash/lodash](https://github.com/lodash/lodash) A modern JavaScript utility library delivering modularity, performance, & extras.