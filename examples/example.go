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
