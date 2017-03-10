package sic

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"

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

// getSubFD returns the subFD of the input
func (c *Cryption) getSubFD() error {
	var err error
	var magic uint16
	binary.Read(c.Input, binary.LittleEndian, &magic) // magic
	k, err := c.Input.ReadByte()                      // sver
	if err != nil {
		return err
	}
	fmt.Println(k)
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
func (c *Cryption) Decrypt() error {
	var err error
	if len(c.Password) == 0 {
		return ErrNoCredentialsSet
	}
	if c.Input == nil {
		return ErrNoInputSet
	}
	if err = c.getSubFD(); err != nil {
		return err
	}
	c.Raw, err = c.getOutput()
	if err != nil {
		return err
	}
	return c.parseDatabase()
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
