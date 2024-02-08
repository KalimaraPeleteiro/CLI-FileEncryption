package criptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(filepath string, password []byte) {
	sourceFile, err := os.Open(filepath)

	if err != nil {
		panic(err.Error())
	}

	defer sourceFile.Close()

	data, err := io.ReadAll(sourceFile)
	if err != nil {
		panic(err.Error())
	}

	key := password
	nonce := make([]byte, 12) // A nonce is an arbitrary number that is used only once in a cryptographic communication.
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	derivedKey := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	blockCipher, err := aes.NewCipher(derivedKey)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		panic(err.Error())
	}

	cipherData := aesgcm.Seal(nil, nonce, data, nil)
	cipherData = append(cipherData, nonce...)

	destinationFile, err := os.Create(filepath)
	if err != nil {
		panic(err.Error())
	}

	defer destinationFile.Close()

	_, err = destinationFile.Write(cipherData)
	if err != nil {
		panic(err.Error())
	}
}
