package authentication

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

type Auth struct {
	SSL     bool
	BaseUrl string
}

func (auth Auth) Authenticate(username, pswd string) error {
	_, err := GenSessionID()
	if err != nil {
		return errors.Wrap(err, "Error generating session ID")
	}

	return nil
}

// GenSessionID generates a session ID in Hexadecimal format
// with 32 characters. The session ID is generated with 256 random bytes
// that are then encrypted using MD5 encryption method.
func GenSessionID() (string, error) {
	b := make([]byte, 256)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.Wrap(err, "Error reading random byte for session ID generation")
	}

	md5Sess := md5.New()
	io.WriteString(md5Sess, string(b))

	checkSumHex := fmt.Sprintf("%x", md5Sess.Sum(nil))
	return strings.ToUpper(checkSumHex), nil
}
