package sic

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestDecrytionSuccess(t *testing.T) {
	file, err := os.Open("decrypt_test.db")
	if err != nil {
		t.Errorf("could not read file: %v", err)
	}
	raw, err := Decrypt(file, "foobar")
	if err != nil {
		t.Errorf("could not decrypt: %v", err)
	}
	xmlFile, err := ioutil.ReadFile("decrypt_test.xml")
	if err != nil {
		t.Errorf("could not read xml file: %v", err)
	}
	if bytes.Compare(xmlFile, raw) != 0 {
		t.Errorf("file content diffs from expected one")
	}
}

func TestDecryptionInvalidPassword(t *testing.T) {
	file, err := os.Open("decrypt_test.db")
	if err != nil {
		t.Errorf("could not read file: %v", err)
	}
	if _, err = Decrypt(file, "definetly not correct"); err != ErrIncorrectPassword {
		t.Errorf("could not decrypt: %v", err)
	}
}
