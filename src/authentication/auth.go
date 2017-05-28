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
	TypeSession          = "Session"
	SubTypeInitSession   = "InitSession"
	SubTypeCheckPassword = "CheckPassword"

	FailMsg    = "FAIL"
	SuccessMsg = "WINFAIL"
)

// Session is a struct that holds the session ID
// as well as the number used for encryption
type Session struct {
	SessionID string
	EncNum    int64
}

// Auth is a struct that holds information about
// the connection. It also provide a set of functions that
// are used for authentication, and session management
type Auth struct {
	SSL     bool
	BaseUrl string
	Path    string
}

func (auth Auth) Authenticate(username, pswd string) (*Session, error) {
	session, err := auth.NewSession()
	if err != nil {
		return nil, errors.Wrap(err, "Error creating the new session.")
	}

	// Check user credentials
	err = auth.checkCredentials(username, pswd, session)
	if err != nil {
		return nil, errors.Wrap(err, "Authentication failed.")
	}

	return session, nil
}

func (auth Auth) NewSession() (*Session, error) {
	// Generate the new session ID
	sessionID, err := GenSessionID()
	if err != nil {
		return nil, errors.Wrap(err, "Error generating session ID.")
	}

	// Initialise the session and retrieve the encryption number from server
	sessionNum, err := auth.initSession(sessionID)
	if err != nil {
		return nil, errors.Wrap(err, "Could not initialise the session.")
	}

	return &Session{
		SessionID: sessionID,
		EncNum:    sessionNum,
	}, nil
}

// getHttp returns the http protocol for url generation
// http if auth.SSL is false, https if true
// Example output "http:" or "https:"
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
