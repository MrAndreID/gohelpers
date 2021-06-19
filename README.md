# MrAndreID / Go Helpers

[![Go Reference](https://pkg.go.dev/badge/github.com/MrAndreID/gohelpers.svg)](https://pkg.go.dev/github.com/MrAndreID/gohelpers) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The `MrAndreID/GoHelpers` package is a collection of functions in the go language.

---

## Table of Contents

* [Install](#install)
* [Usage](#usage)
* [Full Example](#full-example)
* [Versioning](#versioning)
* [Authors](#authors)
* [License](#license)
* [Official Documentation for Go Language](#official-documentation-for-go-language)
* [More](#more)

---

## Install

To use The `MrAndreID/GoHelpers` package, you must follow the steps below:

```sh
go get -u github.com/MrAndreID/gohelpers
```

## Usage

### Create an Error Message

```go
gohelpers.ErrorMessage("error loading the .env file", ".env file not found.")
```

Output:

```sh
2021-02-15 18:45:18 [ ERROR ] Message : error loading the .env file.
2021-02-15 18:45:18 [ ERROR ] Detail : .env file not found.
```

### JSON Encode

```go
fmt.Println(gohelpers.JSONEncode(map[string]interface{}{"First Name": "Andrea", "Last Name": "Adam"}))
```

Output:

```sh
{"First Name":"Andrea","Last Name":"Adam"}
```

### Generate a Bytes

```go
fmt.Println(gohelpers.Bytes(4))
```

Output:

```sh
[95 113 200 231]
```

### Generate a Random Bytes

```go
fmt.Println(gohelpers.RandomByte(32))
```

Output:

```sh
1UlrTYbNJioQPyBEKpV5BFtgqV6t5fEvjSaO8ApGRHs=
```

### Generate a Random Strings

```go
fmt.Println(gohelpers.Random("str", 10))
```

Output:

```sh
XBMUH3qvXh
```

### Generate a Random Int

```go
fmt.Println(gohelpers.Random("int", 4))
```

Output:

```sh
6111
```

### Generate Key

```go
fmt.Println(gohelpers.GenerateKey(32))
```

Output:

```sh
7f2f9d692d200e20133428c832b80f8e21702437fcd28ba2ac8c5aaa3a978b2d
```

### Encrypt

```go
key := gohelpers.GenerateKey(32)

encryptedData, err := gohelpers.Encrypt(key, "Andrea Adam")
if err != nil {
    gohelpers.ErrorMessage("something went wrong when encrypting data", err)
}

fmt.Println(encryptedData)
```

Output:

```sh
b9ab3d8bde4092791b50142be86dfdc70688d81f42fa4aa06c88bcb1af6dfaa4f6c920ec157874
```

### Decrypt

```go
key := gohelpers.GenerateKey(32)

plainText, err := gohelpers.Decrypt(key, "b9ab3d8bde4092791b50142be86dfdc70688d81f42fa4aa06c88bcb1af6dfaa4f6c920ec157874")
if err != nil {
    gohelpers.ErrorMessage("something went wrong when decrypting data", err)
}

fmt.Println(plainText)
```

Output:

```sh
Andrea Adam
```

### Get New Line

```go
fmt.Println("Andrea" + gohelpers.GetNewLine() + "Adam")
```

Output:

```sh
Andrea
Adam
```

### Merge Maps

```go
map1 := map[string]interface{}{"FirstName": "Andrea", "LastName": "Adam"}
map2 := map[string]interface{}{"Age": 21}
map3 := map[string]interface{}{"FirstName": "Andrea", "MidName": nil, "LastName": "Adam"}

fmt.Println(gohelpers.MergeMaps(map1, map2, map3))
```

Output:

```sh
map[Age:21 FirstName:Andrea LastName:Adam MidName:<nil>]
```

### Generate Encrypted Key

```go
key := gohelpers.GenerateKey(32)

encryptedKey, err := gohelpers.GenerateEncryptedKey([]string{"Andrea", "Adam"}, "_", key)
if err != nil {
    gohelpers.ErrorMessage("something went wrong when generating encrypted key", err)
}

fmt.Println(encryptedKey)
```

Output:

```sh
ccaa9be63b4699a53166e1cf4a0086ff3ced25dca2f0672b9cb22309f5270087e7947cb643e579
```

### Generate Encrypted Key With Datetime

```go
key := gohelpers.GenerateKey(32)

encryptedKey, err := gohelpers.GenerateEncryptedKeyWithDatetime([]string{"Andrea", "Adam"}, "_", key, time.Now())
if err != nil {
    gohelpers.ErrorMessage("something went wrong when generating encrypted key with datetime", err)
}

fmt.Println(encryptedKey)
```

Output:

```sh
4e9206fc51372eb6f983a06abdfdb23e7cff0cf32d8f418997428547ed6aef438274fa6054e9d63c96c9c929a6da4c2268700da5d0fe3f06f348a3
```

### Ungenerate Encrypted Key

```go
key := gohelpers.GenerateKey(32)

data, err := gohelpers.UngenerateEncryptedKey("ccaa9be63b4699a53166e1cf4a0086ff3ced25dca2f0672b9cb22309f5270087e7947cb643e579", "_", key)
if err != nil {
    gohelpers.ErrorMessage("something went wrong when ungenerating encrypted key", err)
}

fmt.Println(data)
```

Output:

```sh
[Andrea Adam]
```

### Generate Hash And Salt

```go
key := gohelpers.GenerateKey(32)

encryptedHash, encryptedSalt, err := gohelpers.GenerateHashAndSalt("Andrea Adam", 32, key, 5)
if err != nil {
    gohelpers.ErrorMessage("something went wrong when generating hash and salt", err)
}

fmt.Println("Hash : " + encryptedHash + ", Salt : " + encryptedSalt)
```

Output:

```sh
Hash : 18f65095a8a2ad99851072aee8801c73eea67e2fb18866e8b96aaf4fdd996a879ee1a7987dbd2e9cf6803de30b8224eec77c63e0b9fac91e8c36b1c7fbe54589bd28bec89c774d3f6b8ea7b411d6edd8ef07630cf9689e4b, Salt : 2b4e42cf347c74ff1b1e28b99a3d50102313b022b5eac24ba1ba1287c3913c4d878cc94ae3def9908f6574e9e2d78777c993dc147dc93a7e13eccd9ecc418e4205acb763bc623693
```

### Verify Hash And Salt

```go
key := gohelpers.GenerateKey(32)

encryptedHash, encryptedSalt, err := gohelpers.GenerateHashAndSalt("Andrea Adam", 32, key, 5)
if err != nil {
    gohelpers.ErrorMessage("something went wrong when generating hash and salt", err)
}

fmt.Println(gohelpers.VerifyHashAndSalt("Andrea Adam", encryptedHash, encryptedSalt, key))
```

Output:

```sh
true
```

## Full Example

Full Example can be found on the [Go Playground website](https://play.golang.com/p/1kjlYIetAPb).

## Versioning

I use [SemVer](https://semver.org/) for versioning. For the versions available, see the tags on this repository. 

## Authors

**Andrea Adam** - [MrAndreID](https://github.com/MrAndreID/)

## License

MIT licensed. See the LICENSE file for details.

## Official Documentation for Go Language

Documentation for Go Language can be found on the [Go Language website](https://golang.org/doc/).

## More

Documentation can be found [on https://go.dev/](https://pkg.go.dev/github.com/MrAndreID/gohelpers).
