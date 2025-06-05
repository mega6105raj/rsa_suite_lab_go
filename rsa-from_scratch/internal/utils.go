/* Documenation Section:
	Date : 5th June 2025
	Description : This file includes various utility functions that are used acroos the project
			1. StringToBigInt -> Converts a string to a big.Int (for encryption)
			2. BigIntToString -> Converts a big.Int to string (after decryption)
*/

package internal

import (
	"math/big"
)

// This function converts a string to a big.Int (for encryption)
func StringToBigInt(message string) *big.Int {
	return new(big.Int).SetBytes([]byte(message))
}
// This function converts a big.Int to string (after decryption)
func BigIntToString(b *big.Int) string {
	return string(b.Bytes())
}

