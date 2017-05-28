package encryption

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSha1Hex(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected string
	}{
		{"clubsandwich", "918b336294af220d89001fc9dea52e3e640900f6"},
		{"T$(LCy gbhn3wb798", "41c8d2281f04299ed0a8e57027217f8fd49bd91d"},
		{"Elastica", "b3b6405129b41ba8a40fa27d98ce436ff65d1788"},
		{"Elisa123_123_8383:", "7e8d1d441fc5f5661a772335ae3431865070376c"},
		{"Kang0_!###!@123", "5478848f45aded82d52a336ddc7cbc537f3652f5"},
	}

	for _, testCase := range testCases {
		actual := Sha1Hex(testCase.Input)
		assert.Equal(t, testCase.Expected, actual)
	}
}
