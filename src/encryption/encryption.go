package encryption

import (
	"crypto"
	"fmt"
	"io"
)

func Sha1Hex(input string) string {
	h := crypto.SHA1.New()
	io.WriteString(h, input)

	hexStr := fmt.Sprintf("%x", h.Sum(nil))
	return hexStr
}


