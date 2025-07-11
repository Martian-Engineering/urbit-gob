# urbit-gob
### A go implementation of [urbit-ob](https://github.com/urbit/urbit-ob)

---

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/deelawn/urbit-gob)
[![GoReportCard](https://goreportcard.com/badge/github.com/nanomsg/mangos)](https://goreportcard.com/report/github.com/deelawn/urbit-gob)

#### Command line use
```
> go run cmd/main.go patp 0
~zod
> go run cmd/main.go clan ~marzod
star
> go run cmd/main.go --help
Usage: ...main COMMAND args...

Valid commands:

    patp                : converts a number to a @p-encoded string

    patp2dec            : converts a @p-encoded string to a decimal-encoded string

    patp2hex            : converts a @p-encoded string to a hex-encoded string

    patp2point          : converts a @p-encoded string to a point-encoded string

    patq                : converts a number to a @q-encoded string

    patq2dec            : converts a @q-encoded string to a decimal-encoded string

    patq2hex            : converts a @q-encoded string to a hex-encoded string

    patq2point          : converts a @q-encoded string to a point-encoded string

    point2patp          : converts a point-encoded string to a @p-encoded string

    point2patq          : converts a point-encoded string to a @q-encoded string

    hex2patp            : converts a hex-encoded string to a @p-encoded string

    hex2patq            : converts a hex-encoded string to a @q-encoded string

    clan                : determines the ship class of a @p value

    clanpoint           : determines the ship class of an int-encoded point

    sein                : determines the parent of a @p value

    seinpoint           : determines the parent of an int-encoded point

    eqpatq              : performs an equality comparison on @q values

    isvalidpat          : weakly checks if a string is a valid @p or @q value

    isvalidpatp         : validates a @p string

    isvalidpatq         : validates a @q string
```

#### Module use
```go
package main

import "github.com/deelawn/urbit-gob/co"

func main() {

	// name = ~fipfes
	name, err := co.Patp("65535")
	if err != nil {
		panic(err)
	}

	// sponsor = ~fes
	sponsor, err := co.Sein(name)
	if err != nil {
		panic(err)
	}
}
```
