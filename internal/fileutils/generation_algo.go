package fileutils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes     = "0123456789"
)

// generatePassword creates a random password with the specified length
// and includes letters, special characters, and/or numbers based on the flags.
func GeneratePassword(length int, useLetters, useSpecial, useNum bool) (string, error) {
	var chars string
	if useLetters {
		chars += letterBytes
	}
	if useSpecial {
		chars += specialBytes
	}
	if useNum {
		chars += numBytes
	}

	if len(chars) == 0 {
		return "", fmt.Errorf("no character sets selected")
	}

	b := make([]byte, length)
	for i := range b {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		b[i] = chars[idx.Int64()]
	}

	// Ensure at least one character from each selected set is included
	if useLetters {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		b[0] = letterBytes[idx.Int64()]
	}
	if useSpecial {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(specialBytes))))
		if err != nil {
			return "", err
		}
		b[1] = specialBytes[idx.Int64()]
	}
	if useNum {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(numBytes))))
		if err != nil {
			return "", err
		}
		b[2] = numBytes[idx.Int64()]
	}

	// Shuffle the slice to avoid predictable patterns
	for i := range b {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			return "", err
		}
		b[i], b[j.Int64()] = b[j.Int64()], b[i]
	}

	return string(b), nil
}
