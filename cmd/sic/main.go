package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"syscall"

	sic "github.com/clebs/golang-safe-in-cloud"
	"golang.org/x/term"
)

func main() {
	db := flag.String("f", "SafeInCloud.db", "SafeInCloud Database file to decrypt")
	out := flag.String("o", "decrypted.txt", "Output file to store decrypted data (default decrypted.txt)")
	flag.Parse()

	file, err := os.Open(*db)
	if err != nil {
		log.Fatalf("could not read file: %s", err)
	}

	fmt.Print("Please provide the master password:")
	pw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalf("Could not read password: %s", err)
	}

	raw, err := sic.Decrypt(file, string(pw))
	if err != nil {
		log.Fatalf("could not decrypt: %s", err)
	}
	// if err := write(*out, raw); err != nil {
	// 	log.Fatalf("Could not write raw data to %s: %s", *out, err)
	// }

	d, err := sic.Unmarshal(raw)
	if err != nil {
		log.Printf("could not process database, saving raw output: %s", err)
		if err := write(*out, raw); err != nil {
			log.Fatalf("Could not write raw data to %s: %s", *out, err)
		}
	}

	json, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		log.Fatalf("Could not format contents: %s", err)
	}

	if err := write(*out, json); err != nil {
		log.Fatalf("Could not write data to %s: %s", *out, err)
	}
	fmt.Printf("\nDecrypted data saved to: %s!\n", *out)
}

func write(file string, data []byte) error {
	return ioutil.WriteFile(file, data, 0777)
}
