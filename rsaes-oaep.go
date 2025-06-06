/* Implements RSAES-OAEP variant */
package main

import (
	"bufio"
	"os"
	"strings"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKey := &privateKey.PublicKey

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a message to encrypt ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	message := []byte(strings.TrimSpace(userInput))

	label := []byte("")
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, message, label)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nEncrypted message : %x\n", ciphertext)

	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, label)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nDecrypted message : %s\n", plaintext)
}
