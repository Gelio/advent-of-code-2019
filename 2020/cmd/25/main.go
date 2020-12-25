package main

import "fmt"

func main() {
	// Input
	cardPublicKey := 8335663
	doorPublicKey := 8614349

	cardLoopSize := getLoopSize(cardPublicKey)
	fmt.Println("Result A:", getEncryptionKey(doorPublicKey, cardLoopSize))
}
