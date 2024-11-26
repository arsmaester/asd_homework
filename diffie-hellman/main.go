package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generateRandomBigInt(max *big.Int) (*big.Int, error) {
	return rand.Int(rand.Reader, max)
}

func main() {
	P := big.NewInt(1231)
	G := big.NewInt(5)

	fmt.Println("Publicly known values:")
	fmt.Printf("P (prime): %d, G (generator): %d\n", P, G)

	alicePrivateKey, _ := generateRandomBigInt(P)
	bobPrivateKey, _ := generateRandomBigInt(P)

	fmt.Println("\nPrivate keys (kept secret):")
	fmt.Printf("Alice's Private Key: %d\n", alicePrivateKey)
	fmt.Printf("Bob's Private Key: %d\n", bobPrivateKey)

	alicePublicKey := new(big.Int).Exp(G, alicePrivateKey, P)
	bobPublicKey := new(big.Int).Exp(G, bobPrivateKey, P)

	fmt.Println("\nPublic keys (shared publicly):")
	fmt.Printf("Alice's Public Key: %d\n", alicePublicKey)
	fmt.Printf("Bob's Public Key: %d\n", bobPublicKey)

	aliceSharedSecret := new(big.Int).Exp(bobPublicKey, alicePrivateKey, P)
	bobSharedSecret := new(big.Int).Exp(alicePublicKey, bobPrivateKey, P)

	fmt.Println("\nShared secrets (computed independently):")
	fmt.Printf("Alice's Shared Secret: %d\n", aliceSharedSecret)
	fmt.Printf("Bob's Shared Secret: %d\n", bobSharedSecret)

	if aliceSharedSecret.Cmp(bobSharedSecret) == 0 {
		fmt.Println("\nKey exchange successful! Both parties have the same shared secret.")
	} else {
		fmt.Println("\nKey exchange failed! Shared secrets do not match.")
	}
}
