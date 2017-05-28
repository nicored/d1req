package encryption

import (
	"crypto"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
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
	binary.LittleEndian.PutUint32(numBytes, uint32(num))

	output := make([]byte, len(input))

	for i, _ := range input {
		output[i] = input[i] ^ numBytes[i%len(numBytes)]
	}

	return strings.ToUpper(hex.EncodeToString(output)), nil
}
