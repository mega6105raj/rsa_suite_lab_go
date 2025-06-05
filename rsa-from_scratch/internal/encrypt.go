/* Documentation Section :
	Date : 5th June 2025
	Description : The following code performs the encryption. It uses only the public key (n,e) and the plain text to be encrypted
		      Built-in libraries are used for modular exponentiation, rather than implementing fast exponentiation algorithm from scratch */

package internal

import (
	"errors"
	"math/big"
)

//The following function encrypts the plaintext using the public key (n,e) and return the cipher text
func EncryptRSA(m *big.Int, e,n *big.Int) (*big.Int, error) {
	if m.Cmp(n) >= 0 {
		return nil, errors.New("\nMessage must be smaller than modulus n")
	}

	c := new(big.Int).Exp(m, e, n)
	return c, nil
}

