package criptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Decrypt(filepath string, password []byte) {
	sourceFile, err := os.Open(filepath)

	if err != nil {
		panic(err.Error())
	}

	defer sourceFile.Close()

	cipherData, err := io.ReadAll(sourceFile)
	if err != nil {
		panic(err.Error())
	}

	key := password
	salt := cipherData[len(cipherData)-12:]
	strSalt := hex.EncodeToString(salt)
	nonce, err := hex.DecodeString(strSalt)
	if err != nil {
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

	originalData, err := aesgcm.Open(nil, nonce, cipherData[:len(cipherData)-12], nil)
	if err != nil {
		panic(err.Error())
	}

	destinationFile, err := os.Create(filepath)
	if err != nil {
		panic(err.Error())
	}

	defer destinationFile.Close()

	_, err = destinationFile.Write(originalData)
	if err != nil {
		panic(err.Error())
	}
}
