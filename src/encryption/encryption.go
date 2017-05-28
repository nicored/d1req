package encryption

import (
	"crypto"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
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
	numBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(numBytes, uint32(num))

	byteAt := len(numBytes)
	retVal := ""
	for _, char := range input {
		if byteAt == 0 {
			byteAt = len(numBytes) - 1
		} else {
			byteAt -= 1
		}
		numByte := numBytes[byteAt]

		// Xor operation on char code and current byte in num
		xorInt := int64(char) ^ int64(numByte)

		// Convert xor result to hexadecimal formal
		// and append '0' if xorHex <= F
		xorHex := strconv.FormatInt(xorInt, 16)
		if len(xorHex) == 1 {
			retVal += "0" + xorHex
		} else {
			retVal += xorHex
		}
	}

	return strings.ToUpper(retVal), nil
}
