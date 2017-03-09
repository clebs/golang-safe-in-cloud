package sic

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

var (
	// ErrIncorrectCredentials means that the credentials are incorrect
	ErrIncorrectCredentials = errors.New("Incorrect credentials")
	// ErrNoCredentialsSet means that the credentials are not set
	ErrNoCredentialsSet = errors.New("No credentials set. Forgotten to set a password?")
	// ErrNoInputSet means that no Input was set
	ErrNoInputSet = errors.New("No Input was set. Forgotten to set a Input?")
)

// Cryption is the struct which holds the
// de- en-cryption stuff
type Cryption struct {
	Input          *bufio.Reader
	Password       string
	CryptoGetSubFD cryptContainer
	Crypto         cryptContainer
	SubFD          *bufio.Reader
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

// getSubFD returns the subFD of the input
func (c *Cryption) getSubFD() error {
	var err error
	var magic uint16
	binary.Read(c.Input, binary.LittleEndian, &magic) // magic
	_, err = c.Input.ReadByte()                       // sver
	if err != nil {
		return err
	}
	c.CryptoGetSubFD.Salt, err = c.readByteArray()
	if err != nil {
		return err
	}
	c.CryptoGetSubFD.IV, err = c.readByteArray() // nonce
	if err != nil {
		return err
	}
	block, err := aes.NewCipher(pbkdf2.Key([]byte(c.Password), c.CryptoGetSubFD.Salt, 10000, 32, sha1.New))
	if err != nil {
		return err
	}
	mode := cipher.NewCBCDecrypter(block, c.CryptoGetSubFD.IV)
	c.Crypto.Salt, err = c.readByteArray()
	if err != nil {
		return err
	}
	src, err := c.readByteArray()
	if err != nil {
		return err
	}
	mode.CryptBlocks(src, src)

	c.SubFD = bufio.NewReader(bytes.NewBuffer(src))

	return nil
}

// getOutput returns the decrypted output from the database
// it have to be a given subFD before
func (c *Cryption) getOutput() ([]byte, error) {
	var err error
	c.Crypto.IV, err = c.readByteArraySubFD()
	if err != nil {
		return nil, err
	}
	c.Crypto.Password, err = c.readByteArraySubFD()
	if err != nil {
		return nil, ErrIncorrectCredentials
	}
	_, err = c.readByteArraySubFD() // check
	if err != nil {
		return nil, err
	}
	pbkdf2.Key(c.Crypto.Password, c.Crypto.Salt, 1000, 32, sha1.New)
	block, err := aes.NewCipher(c.Crypto.Password)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(c.Input)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, c.Crypto.IV)
	mode.CryptBlocks(data, data)
	zReader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	output, err := ioutil.ReadAll(zReader)
	if err != nil {
		return nil, err
	}
	zReader.Close()
	return output, nil
}

// Decrypt is the Endpoint to start decrypting of
// the database. Before that you need to set password
func (c *Cryption) Decrypt() (*Database, error) {
	var err error
	if len(c.Password) == 0 {
		return nil, ErrNoCredentialsSet
	}
	if c.Input == nil {
		return nil, ErrNoInputSet
	}
	if err = c.getSubFD(); err != nil {
		return nil, err
	}
	output, err := c.getOutput()
	if err != nil {
		return nil, err
	}
	var db Database
	if err = xml.Unmarshal(output, &db); err != nil {
		log.Fatal(err)
	}
	return &db, nil
}

// readByteArraySubFD reads a byte array with the size of
// the given size in the first byte from the subFD
func (c *Cryption) readByteArraySubFD() ([]byte, error) {
	size, err := c.SubFD.ReadByte()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, size)
	if err = binary.Read(c.SubFD, binary.LittleEndian, &buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// readByteArray reads a byte array with the given size
// of the first byte
func (c *Cryption) readByteArray() ([]byte, error) {
	size, err := c.Input.ReadByte()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, size)
	if err = binary.Read(c.Input, binary.LittleEndian, &buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// SetPassword sets the password for the basic
// auth
func (c *Cryption) SetPassword(pw string) {
	c.Password = pw
}
