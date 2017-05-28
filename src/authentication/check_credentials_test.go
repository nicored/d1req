package authentication

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"
)

func TestCheckCredentialsSuccess(t *testing.T) {
	session := &Session{
		SessionID: sessionID,
		EncNum:    89,
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(auth), "SendRequest", func(_ Auth, _ url.URL) ([]byte, error) {
		return []byte("WINFAIL"), nil
	})

	err := auth.checkCredentials("nico", "myPassword", session)
	assert.NoError(t, err)
}

func TestCheckCredentialsInvalidCredentials(t *testing.T) {
	session := &Session{
		SessionID: sessionID,
		EncNum:    89,
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(auth), "SendRequest", func(_ Auth, _ url.URL) ([]byte, error) {
		return []byte("FAIL"), nil
	})

	err := auth.checkCredentials("nico", "myPassword", session)
	assert.Error(t, err)
	assert.Equal(t, "Invalid credentials.", err.Error())
}

func TestCheckCredentialsTimedOut(t *testing.T) {
	session := &Session{
		SessionID: sessionID,
		EncNum:    89,
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(auth), "SendRequest", func(_ Auth, _ url.URL) ([]byte, error) {
		return []byte(""), nil
	})

	err := auth.checkCredentials("nico", "myPassword", session)
	assert.Error(t, err)
	assert.Equal(t, "Authentication request timed out.", err.Error())
}

func TestEncryptUsername(t *testing.T) {
	// Expected values are based on the output from the javascript version
	// of this program
	testCases := []struct {
		Username string
		EncNum   int64
		Expected string
	}{
		{"clubsandwich", 89, "87AB32B2F89806EEEB0C125FD0513B65153EFC40"},
		{"T$(LCy gbhn3wb798", 11, "0C05FCF33B548D7635C364FA45012C865B4E03C5"},
		{"Elastica", 12, "8EAC69C7BD6668507EF56663CA6DE68D6FB2A4BB"},
		{"Elisa123_123_8383:", 99, "8B6551720AF7286F3E36DFF6DB674ECA996AF879"},
		{"Kang0_!###!@123", 69, "02CF395A855CA9CB0ABE6A5176141D3653BBBD60"},
	}

	for _, testCase := range testCases {
		encr, err := EncryptUsername(testCase.Username, testCase.EncNum)
		assert.NoError(t, err)
		assert.Equal(t, testCase.Expected, encr)
	}
}

func TestEncryptPassword(t *testing.T) {
	// Expected values are based on the output from the javascript version
	// of this program
	testCases := []struct {
		Username string
		EncNum   int64
		Expected string
	}{
		{"clubsandwich", 89, "4ACB1CB14160075F2BF515F934A15423E1630044"},
		{"T$(LCy gbhn3wb798", 11, "453616458A4DA6EB7FB168B7053EF5ABE639E51D"},
		{"Elastica", 12, "B4C906225EB1269CE669C9868502CA03D4F58B3D"},
		{"Elisa123_123_8383:", 99, "C93B5DD058C2E74CA870C2051F327706E7D6B7F7"},
		{"Kang0_!###!@123", 69, "6252271F8A565E462C454EBCBBCCA559348FDD22"},
	}

	for _, testCase := range testCases {
		encr, err := EncryptPassword(testCase.Username, testCase.EncNum)
		assert.NoError(t, err)
		assert.Equal(t, testCase.Expected, encr)
	}
}
