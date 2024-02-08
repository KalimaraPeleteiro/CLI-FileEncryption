package main

import (
	"bytes"
	"os"
)

func isFileValid(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}

func isPasswordValid(password, confirmationPassword []byte) bool {
	if !bytes.Equal(password, confirmationPassword) {
		return false
	}

	return true
}
