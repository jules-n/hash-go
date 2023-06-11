package main

import (
	"golang.org/x/crypto/scrypt"
)

func generateScryptHash(password, salt []byte) ([]byte, error) {
	const (
		N      = 16384 // Кількість ітерацій
		R      = 8     // Вимоги пам'яті
		P      = 1     // Паралелізм
		KeyLen = 32    // Довжина вихідного хешу
	)

	hash, err := scrypt.Key(password, salt, N, R, P, KeyLen)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
