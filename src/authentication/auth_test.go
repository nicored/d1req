package authentication

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth_NewSession(t *testing.T) {
	session, err := auth.NewSession()
	assert.NoError(t, err)

	assert.Regexp(t, "\\d+", session.EncNum)
}

func TestGenSessionID(t *testing.T) {
	for i := 0; i < 15; i++ {
		sessionID, err := GenSessionID()
		assert.NoError(t, err)

		// Assert that the length of the session ID is 32
		assert.Len(t, sessionID, 32)

		// Assert that output is a proper hexadecimal
		sessionIDBytes := []byte(sessionID)
		dst := make([]byte, hex.DecodedLen(len(sessionIDBytes)))

		_, err = hex.Decode(dst, sessionIDBytes)
		assert.NoError(t, err)
	}
}

func TestGetHttp(t *testing.T) {
	auths := Auth{SSL: true}

	prot := auths.getHttp()
	assert.Equal(t, "https:", prot)

	auths.SSL = false
	prot = auths.getHttp()
	assert.Equal(t, "http:", prot)
}
