package sic

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

// Encrypt is the Endpoint to start encrypting of
// the database. You can get the content via func getRaw()
// via func writeToFile(filename)
func (c *Cryption) Encrypt() error {
	var err error
	c.Raw, err = xml.MarshalIndent(c.DB, " ", "    ")
	if err != nil {
		return err
	}
	if err := c.encrypt(); err != nil {
		return err
	}
	if err := c.saveFile(); err != nil {
		return err
	}
	return nil
}

func (c *Cryption) encrypt() error {
	data := c.addPadding(c.Raw)
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(pbkdf2.Key([]byte(c.Password), c.CryptoGetSubFD.Salt, 10000, 32, sha1.New))
	if err != nil {
		return err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCBCEncrypter(block, iv)
	stream.CryptBlocks(ciphertext[aes.BlockSize:], data)
	f, err := os.Create("a_aes.txt")
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		return err
	}
	fmt.Println("done")
	return nil
}

func (c *Cryption) saveFile() error {

	return nil
}

func (c *Cryption) addPadding(value []byte) []byte {
	var out string
	pad := aes.BlockSize - (len(value) % aes.BlockSize)
	for i := 0; i < pad; i++ {
		out += string(pad)
	}
	fmt.Println(len(out))
	return append(value, []byte(out)...)
}
