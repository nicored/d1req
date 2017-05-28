package encryption

import (
	"crypto"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"strconv"
	"strings"
)

// Sha1Hex encrypt the input using SHA-1 encryption method, and
// returns it in Hexadecimal format in lowercase
func Sha1Hex(input string) (string, error) {
	h := crypto.SHA1.New()

	_, err := io.WriteString(h, input)
	if err != nil {
		return "", errors.Wrap(err, "Errored encrypting to SHA-1")
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// Xor performs bit operations to encrypt the input.
// input is the string to be encrypted
// num is the number used to perform xor comparisons with the input
// The output of Xor is a string representing the result of the operation in Hexadecimal format
// in UPPERCASE
func Xor(input string, num int64) (string, error) {
	// Convert num to binary format
	numBin := fmt.Sprintf("%032s", strconv.FormatInt(num, 2))

	startPos := len(numBin)
	retVal := ""
	for _, char := range input {
		if startPos == 0 {
			startPos = len(numBin) - 8
		} else {
			startPos = startPos - 8
		}

		// Convert 1 byte at a time to int64 format
		comp, err := strconv.ParseInt(numBin[startPos:startPos+8], 2, 64)
		if err != nil {
			return "", errors.Wrap(err, "Error converting byte to int.")
		}

		// Xor operation on char code and current byte in num
		xorInt := int64(char) ^ comp

		// Convert xor result to hexadecimal formal
		// and append '0' if xorHex <= F
		xorHex := strconv.FormatInt(xorInt, 16)
		if len(xorHex) == 1 {
			xorHex = "0" + xorHex
		}

		// Append to the chain
		retVal += xorHex
	}

	return strings.ToUpper(retVal), nil
}
