package authentication

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"
)

var sessionID = "FCD21025D7ABE86FAE748E4C7A6507F1"
var auth = Auth{
	BaseUrl: "test.div1.io",
	SSL:     true,
	Path:    "/test/1/exec",
}

type causer interface {
	Cause() error
}

func TestAuth_initSessionSuccess(t *testing.T) {
	monkey.PatchInstanceMethod(reflect.TypeOf(auth), "SendRequest", func(_ Auth, _ url.URL) ([]byte, error) {
		return []byte("69"), nil
	})

	rNum, err := auth.initSession(sessionID)
	assert.NoError(t, err)
	assert.Equal(t, int64(69), rNum)
}

func TestAuth_initSessionErrorServerFail(t *testing.T) {
	monkey.PatchInstanceMethod(reflect.TypeOf(auth), "SendRequest", func(_ Auth, _ url.URL) ([]byte, error) {
		return []byte("FAIL"), nil
	})

	_, err := auth.initSession(sessionID)
	assert.Error(t, err)
	assert.Equal(t, "Server did not accept the initial request.", err.Error())
}

func TestAuth_initSessionErrorRespIsNotANumber(t *testing.T) {
	monkey.PatchInstanceMethod(reflect.TypeOf(auth), "SendRequest", func(_ Auth, _ url.URL) ([]byte, error) {
		return []byte("ABC"), nil
	})

	_, err := auth.initSession(sessionID)
	assert.Error(t, err)

	errorMsg := "Error occurred handling the response during session initialisation."
	assert.Equal(t, errorMsg, err.Error()[:len(errorMsg)])
}
