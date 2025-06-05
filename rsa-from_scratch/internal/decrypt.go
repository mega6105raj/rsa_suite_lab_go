/* Documentaton Section :
	Date : 5th June 2025
	Description : This file contains the decryption logic */

package internal

import (
	"math/big"
)

//This function decrypts the ciphertext using the private key d
func DecryptRSA(c, d, n *big.Int) *big.Int {
	m := new(big.Int).Exp(c, d, n)
	return m
}
