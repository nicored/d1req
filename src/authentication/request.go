package authentication

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func (auth Auth) SendRequest(urlCnf url.URL) ([]byte, error) {
	urlStr := auth.getHttp() + urlCnf.String()

	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, errors.Wrap(err, "An error occurred sending the request")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "An error occurred reading the response body from server")
	}

	return body, nil
}
