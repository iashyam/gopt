package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	input := "This is super secret"
	hashed_input := hash(input)
	fmt.Println("hello world!")
	fmt.Printf("Input: %s\n hash:%s\n ", input, hashed_input)
}

func hash(input string) string {

	inputInBytes := []byte(input)
	hash := sha256.Sum256(inputInBytes)

	return hex.EncodeToString(hash[:])
}
