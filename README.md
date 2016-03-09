**ATTENTION: THIS PACKAGE UNDER DEVELOPMENT**

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

**Example**

```go
var validEmail = "jhon@example.com"
var invalidEmail = "blabla"

godash.IsEmail(validEmail) // result: true
godash.IsEmail(invalidEmail) // result: false
```

Full list of Is* functions;

```go

```

## Thanks & Authors

I use code/got inspiration from these excellent libraries:

- [asaskevich/govalidator](https://github.com/asaskevich/govalidator) [Go] Package of validators and sanitizers for strings, numerics, slices and structs
- [lodash/lodash](https://github.com/lodash/lodash) A modern JavaScript utility library delivering modularity, performance, & extras.