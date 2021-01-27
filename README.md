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
2009/11/10 23:00:00 -------------------- Start Error Message --------------------
2009/11/10 23:00:00 Message => error loading the .env file.
2009/11/10 23:00:00 Error =>  .env file not found.
2009/11/10 23:00:00 -------------------- End Of Error Message --------------------
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

### Handle JSON Response

```go
fmt.Println(gohelpers.HandleJSONResponse("success", "valid data", nil))
```

Output:

```sh
2009/11/10 23:00:00 Closing
2009/11/10 23:00:00 
{"status":"success","message":"valid data","data":null}
```

## Full Example

Full Example can be found on the [Go Playground website](https://play.golang.com/p/g_FZ4HkB6yY).

## Versioning

I use [SemVer](https://semver.org/) for versioning. For the versions available, see the tags on this repository. 

## Authors

**Andrea Adam** - [MrAndreID](https://github.com/MrAndreID/)

## License

MIT licensed. See the LICENSE file for details.

## Official Documentation for Go Language

Documentation for Go Language can be found on the [Go Language website](https://golang.org/doc/).

## More

Documentation can be found [on godoc.org](https://godoc.org/github.com/MrAndreID/gohelpers).
