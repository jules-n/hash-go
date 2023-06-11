package main

import (
	"fmt"
)

func main() {

	passPhrase := "i don't get sleep"

	saltPhrase := "kings of cringe"

	password := calculateSHA256(passPhrase)

	hash, err := generateScryptHash([]byte(password), []byte(saltPhrase))
	if err != nil {
		fmt.Println("Помилка при генерації scrypt hash:", err)
		return
	}

	fmt.Println("SHA-256 Hash:", password)
	fmt.Printf("Scrypt Hash: %x\n", hash)
}
