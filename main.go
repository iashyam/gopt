package main

import ("fmt"
		"crypto/sha256"
		"encoding/hex"
		)

func main(){
	fmt.Println("hello world");
	input := "This is super secret"
	fmt.Printf("Input: %s\n hash:%s\n ", input, hash(input))
}


func hash(input string) string{
	
	inputInBytes := []byte(input);
	hash := sha256.Sum256(inputInBytes)

	return hex.EncodeToString(hash[:])
}
