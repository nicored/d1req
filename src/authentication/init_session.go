package authentication

import (
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// InitSession sends a request to initialise a new session, and retrieves
// an encryption number from the server that is used to encrypt data in all requests
func (auth Auth) initSession(sessionID string) (int64, error) {
	urlCnf := url.URL{
		Host:     auth.BaseUrl,
		Path:     auth.Path,
		RawQuery: genInitSessionQuery(sessionID),
	}

	// Sending the request
	resp, err := auth.SendRequest(urlCnf)
	if err != nil {
		return 0, errors.Wrap(err, "Error occurred while initialising the session.")
	}

	// Check whether or not the server accepted to initialise the session based on the request
	respStr := string(resp)
	if respStr == FailMsg {
		return 0, errors.New("Server did not accept the initial request.")
	}

	// We expect the response to be a number, so we convert it
	respNum, err := strconv.Atoi(respStr)
	if err != nil {
		return 0, errors.Wrap(err, "Error occurred handling the response during session initialisation.")
	}

	return int64(respNum), nil
}
