package sic

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"os"
)

// Cryption is the struct which holds the
// de- en-cryption stuff
type Cryption struct {
	Input          *bufio.Reader
	Password       string
	CryptoGetSubFD cryptContainer
	Crypto         cryptContainer
	SubFD          *bufio.Reader
	Raw            []byte
	DB             Database
}

// cryptContainer holds the crypto stuff
// for one crypt operation
type cryptContainer struct {
	Password []byte
	Salt     []byte
	IV       []byte
}

// NewSafeInCloud return the NewSafeInCloud root object
func NewSafeInCloud() *Cryption {
	return &Cryption{}
}

// SetInputDirect sets the direct input per filecontent
func (c *Cryption) SetInputDirect(content []byte) {
	c.Input = bufio.NewReader(bytes.NewBuffer(content))
}

// SetInputFile sets the input by filename
func (c *Cryption) SetInputFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	c.Input = bufio.NewReader(f)
	return nil
}

// SetPassword sets the password for the basic
// auth
func (c *Cryption) SetPassword(pw string) {
	c.Password = pw
}

// GetRawXML returns the plain xml of the decrypted DB
func (c *Cryption) GetRawXML() []byte {
	return c.Raw
}

// SaveXMLToFile saves the plain xml to a file by the given
// name
func (c *Cryption) SaveXMLToFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	if err = f.Truncate(0); err != nil {
		return err
	}
	if _, err = f.Write(c.Raw); err != nil {
		return err
	}
	return err
}

// GetDatabase returns the database struct within the
// unmarshaled xml and their helper functions
func (c *Cryption) parseDatabase() error {
	return xml.Unmarshal(c.Raw, &c.DB)
}
