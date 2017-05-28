package authentication

import (
	"net/url"
	"strings"

	"encryption"
	"github.com/pkg/errors"
)

// checkCredentials encrypts username and password, and sends an auth request
// to authenticate the user against the session
func (auth Auth) checkCredentials(username string, pswd string, session *Session) error {
	encryptedUsername, err := EncryptUsername(username, session.EncNum)
	if err != nil {
		return errors.Wrap(err, "Username encryption failed.")
	}

	encryptedPswd, err := EncryptPassword(pswd, session.EncNum)
	if err != nil {
		return errors.Wrap(err, "Password encryption failed.")
	}

	urlCnf := url.URL{
		Host:     auth.BaseUrl,
		Path:     auth.Path,
		RawQuery: genCheckPswdQuery(encryptedUsername, encryptedPswd, session.SessionID),
	}
	resp, err := auth.SendRequest(urlCnf)
	if err != nil {
		return errors.Wrap(err, "An error occurred during authentication.")
	}

	respStr := string(resp)
	if respStr == FailMsg {
		return errors.New("Invalid credentials.")
	}

	if respStr == SuccessMsg {
		return nil
	}

	return errors.New("Authentication request timed out.")
}

// EncryptUsername encrypts a plain username following a set of rules:
// Encrypt username using Xor operations
// Encrypt result to sha1 and to hexadecimal format
// Return result in UPPERCASE
func EncryptUsername(username string, sessNum int64) (string, error) {
	sessNum += 1

	xorUsername, err := encryption.Xor(username, sessNum)
	if err != nil {
		return "", errors.Wrap(err, "Error encrypting username.")
	}

	hashedXorUsername, err := encryption.Sha1Hex(xorUsername)
	if err != nil {
		return "", errors.Wrap(err, "Error encrypting username.")
	}

	return strings.ToUpper(hashedXorUsername), nil
}

// EncryptPassword encrypts a plain password following a set of rules:
// Encrypt plainPassword to sha1 and to hexadecimal format
// Encrypt the result using Xor operations
// Encrypt the result to sha1 and to hexadecimal format
// Return the result in UPPERCASE
func EncryptPassword(plainPassword string, sessNum int64) (string, error) {
	hashedPswd, err := encryption.Sha1Hex(plainPassword)
	if err != nil {
		return "", errors.Wrap(err, "Error encrypting password.")
	}

	xorHashedPswd, err := encryption.Xor(hashedPswd, sessNum)
	if err != nil {
		return "", errors.Wrap(err, "Error encrypting password.")
	}

	hashedXorHashedPswd, err := encryption.Sha1Hex(xorHashedPswd)
	if err != nil {
		return "", errors.Wrap(err, "Error encrypting password.")
	}

	return strings.ToUpper(hashedXorHashedPswd), nil
}

// genCheckPswdQuery generates the query part of the url with
// type, subtype, encrypted name, encrypted password, and the session id
func genCheckPswdQuery(name, pswd, sessionId string) string {
	queryParams := url.Values{}

	queryParams.Add("Type", TypeSession)
	queryParams.Add("SubType", SubTypeCheckPassword)
	queryParams.Add("Name", name)
	queryParams.Add("Password", pswd)
	queryParams.Add("SessionID", sessionId)

	return "Command&" + queryParams.Encode()
}
