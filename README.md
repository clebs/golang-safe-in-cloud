[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/maxibanki/sicGo/issues)
[![GoDoc](https://godoc.org/github.com/maxibanki/sicGo?status.svg)](http://godoc.org/github.com/maxibanki/sicGo)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](http://opensource.org/licenses/MIT)
[![Go Report](https://img.shields.io/badge/Go_report-A+-brightgreen.svg)](http://goreportcard.com/report/maxibanki/sicGo)

# SafeInCloud Golang Decryption

Integrates the decryption of a SafeInCloud database

# Example

```golang
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/maxibanki/golang-safe-in-cloud"
)

func main() {
    file, err := os.Open("SafeInCloud.db")
    if err != nil {
        log.Fatalf("could not read file: %v", err)
    }
    raw, err := sic.Decrypt(file, "foobar")
    if err != nil {
        log.Fatalf("could not decrypt: %v", err)
    }
    fmt.Println(string(raw))
}
```