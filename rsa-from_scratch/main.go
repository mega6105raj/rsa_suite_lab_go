/* Documentaion Section:
	Date : 5th June 2025
	Description : This project aims to provide rsa implemenation from scratch.
	This file acts as the enrty point to test the RSA functions implemented */

package main

import (
	"fmt"
	"log"
	"rsa-from-scratch/internal"
	"bufio"
	"os"
	"strings"
)

func main() {
	p, err := internal.GeneratePrime(512) //Hardcoring the bits to be 512
	if err != nil {
		log.Fatal(err)
	}

	q, err := internal.GeneratePrime(512)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nPrime p : ", p)
	fmt.Println("\nPrime q : ", q)

	n, phi := internal.ComputeKeys(p,q)
	fmt.Println("\nModulus n:", n)
	fmt.Println("\nPhi(n) : ", phi)

	e := internal.ChoosePublicExponent(phi)
	fmt.Println("\nPublic exponent : ", e)

	d, err := internal.ComputePrivateExponent(e, phi)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nPrivate exponent d:", d)

	fmt.Print("\nEnter a message : ")
	reader := bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	message = strings.TrimSpace(message)
	m := internal.StringToBigInt(message)

	ciphertext, err := internal.EncryptRSA(m, e, n)
	ct := internal.BigIntToString(ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nEncrypted ciphertext as big.Int:", ciphertext)

	fmt.Println("\nEncrypted message :", ct)
	plaintext := internal.DecryptRSA(ciphertext, d, n)
	fmt.Println("\nDecrypted as big.Int:", plaintext)

	decryptedText := internal.BigIntToString(plaintext)
	fmt.Println("\nDecrypted message:", decryptedText)
}
