package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha1Hex(t *testing.T) {
	// Expected values are based on the output from the javascript version
	// of this program
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

func TestXor(t *testing.T) {
	// Expected values are based on the output from the javascript version
	// of this program
	testCases := []struct {
		Input    string
		Num      int
		Expected string
	}{
		{"ClubSandwich", 89, "1A6C75620A616E642E696368"},
		{"T$(LCy gbhn3wb798", 74, "1E24284C0979206728686E333D62373972"},
		{"Elastica", 10, "4F6C61737E696361"},
		{"Elisa123_123_8383:", 11, "4E6C69736A3132335431323354383338383A"},
		{"Kang0_!###!@123", 12, "47616E673C5F21232F2321403D3233"},
	}

	for _, testCase := range testCases {
		actual := Xor(testCase.Input, testCase.Num)
		assert.Equal(t, testCase.Expected, actual)
	}
}
