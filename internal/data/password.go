package data

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GetHash(plaintext string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func VerifyPassword(hash []byte, plaintext string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(plaintext))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
