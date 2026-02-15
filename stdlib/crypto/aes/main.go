package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"log"
)

func main() {
	key, err := createPrivateKey(aes.BlockSize)
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(*key)
	if err != nil {
		log.Fatal(err)
	}

	iv, err := createIV(aes.BlockSize)
	if err != nil {
		log.Fatal(err)
	}

	plainText := []byte("Hello, World!")
	paddedText := paddingText(plainText, aes.BlockSize)

	cipherText := encrypt(&block, iv, &paddedText)
	_, err = fmt.Printf("cipher = %v\n", cipherText)
	if err != nil {
		log.Fatal(err)
	}

	plainText2 := decrypt(&block, iv, cipherText)
	_, err = fmt.Printf("plain = %s\n", *plainText2)
	if err != nil {
		log.Fatal(err)
	}
}

func createPrivateKey(size int) (*[]byte, error) {
	key := make([]byte, size)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return &key, nil
}

func createIV(size int) (*[]byte, error) {
	iv := make([]byte, size)

	_, err := rand.Read(iv)
	if err != nil {
		return nil, err
	}

	return &iv, nil
}

func paddingText(text []byte, size int) []byte {
	const noRemainder = 0

	padSize := size
	if remainder := len(text) % size; remainder != noRemainder {
		padSize = size - remainder
	}

	// PKCS7
	padBytes := bytes.Repeat([]byte{byte(padSize)}, padSize)

	return append(text, padBytes...)
}

func encrypt(block *cipher.Block, iv *[]byte, text *[]byte) *[]byte {
	cipherText := make([]byte, len(*text))

	// CBC
	enc := cipher.NewCBCEncrypter(*block, *iv)
	enc.CryptBlocks(cipherText, *text)

	return &cipherText
}

func decrypt(block *cipher.Block, iv *[]byte, ctext *[]byte) *[]byte {
	paddedText := make([]byte, len(*ctext))

	// CBC
	dec := cipher.NewCBCDecrypter(*block, *iv)
	dec.CryptBlocks(paddedText, *ctext)

	return &paddedText
}
