/* Documentation Section :
	Date : 5th June 2025
	Description : RSA keygen algo involves choosing two large prime numbers p and q.
	The following program uses the crypto/rand that provides cryptographically secure randomness to generate a random probable priume.
	Rerunning the project severaltimes ensures that for each user, different primes are being chosen which makes this impelmentation resistant to common modulus attack.
	It then calculates n and phi(n).
	Having decided phi(n), e which is relatively prime to phi(n) is chosen. Preference to the usual public expoent Fermat's number is provided n this implementation
	Then the modular invers of e under phi(n) is found using the  built-in library, rather than implementing Euclidean's algorithm from scratch */

package internal

import (
	"crypto/rand"
	"math/big"
	"fmt"
)

//This function generates a cryptographically secure random prime number of given bit size
func GeneratePrime(bits int) (*big.Int, error) {
	prime, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return prime, nil
}

/*
	bits -> Desired number of bits for the prime number (e.g., 512, 1024, 2048,..)
	prime -> Big integer (a lrge prime number)
	p,q -> Private
*/

//This function computes n and phi(n)
func ComputeKeys(p, q *big.Int) (*big.Int, *big.Int) {
	n := new(big.Int).Mul(p, q)

	one := big.NewInt(1)
	pMinus1 := new(big.Int).Sub(p, one)
	qMinus1 := new(big.Int).Sub(q, one)
	phi := new(big.Int).Mul(pMinus1, qMinus1)

	return n, phi
}

/*
	n -> Public
	phi -> Private
*/

//This function finds the public exponent e
func ChoosePublicExponent(phi *big.Int) *big.Int {
	e := big.NewInt(65537)
	gcd := new(big.Int).GCD(nil, nil, e, phi)
	if gcd.Cmp(big.NewInt(1)) == 0 {
		return e
	}

	//If 65537 is not relatively prime to e
	e = big.NewInt(3)
	for {
		gcd := new(big.Int).GCD(nil, nil, e, phi)
		if gcd.Cmp(big.NewInt(1)) == 0 {
			return e
		}
		e.Add(e, big.NewInt(2))
	}
}

/*
	e = 65537 caled the Fermat number is a common choice for RSA.
	It has low hamming efficien which makes it efficient for exponentiation.
	It's usually coprime with phi(n). But for safety even if it fails, we check for new values of e
*/

//This function finds the private exponent d
func ComputePrivateExponent(e, phi *big.Int) (*big.Int, error) {
	d := new(big.Int).ModInverse(e, phi)
	if d == nil {
		return nil, fmt.Errorf("e has no modular invers mod phi")
	}
	return d, nil
}
