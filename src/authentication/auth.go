package authentication

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

const (
	FailMsg = "FAIL"
)

type Auth struct {
	SSL     bool
	BaseUrl string
	Path    string
}

func (auth Auth) Authenticate(username, pswd string) error {
	sessionID, err := GenSessionID()
	if err != nil {
		return errors.Wrap(err, "Error generating session ID")
	}

	_, err = auth.initSession(sessionID)
	if err != nil {
		return errors.Wrap(err, "Could not initialise the session")
	}

	return nil
}

func (auth Auth) getHttp() string {
	httpStr := "http"
	if auth.SSL == true {
		return httpStr + "s:"
	}

	return httpStr + ":"
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
