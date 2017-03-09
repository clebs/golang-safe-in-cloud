[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/maxibanki/sicGo/issues)
[![GoDoc](https://godoc.org/github.com/maxibanki/sicGo?status.svg)](http://godoc.org/github.com/maxibanki/sicGo)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](http://opensource.org/licenses/MIT)
[![Go Report](https://img.shields.io/badge/Go_report-A+-brightgreen.svg)](http://goreportcard.com/report/maxibanki/sicGo)

# SafeInCloud Golang lib

`sicGo` is a golang libary for doing reading and writing operations with the SafeInCloud database.

### Features:
- Decrypt SafeInCloud database
- Parse the xml

### TODO:
- Add Encryption
- Add methods for addCard, addFile, del- etc. 

### Example:
```go
package main

import (
	"fmt"
	"log"

	"github.com/maxibanki/SafeInCloud"
)

func main() {
	c := sic.NewSafeInCloud()
	c.SetInputFile("SafeInCloud.db")
	c.SetPassword("foo")
	db, err := c.Decrypt()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range db.Cards {
		if v.Template {
			continue
		}
		fmt.Println("----------CARD----------")
		fmt.Printf("Title: %s\n", v.Title)
		fmt.Printf("Notes %s\n", v.Notes)
		for _, f := range v.Fields {
			fmt.Printf("\tField: %s Type: '%s' Value: '%s'\n", f.Name, f.Type, f.Value)
		}
	}
}
```