package authentication

import "net/url"

const (
	TypeSession          = "Session"
	SubTypeInitSession   = "InitSession"
	SubTypeCheckPassword = "CheckPassword"
)

// genInitSessionQuery generates the query part of the URL (everything after ?)
// The query includes a type, subtype and the sessionID
func genInitSessionQuery(sessionID string) string {
	queryParams := url.Values{}

	queryParams.Add("Type", TypeSession)
	queryParams.Add("SubType", SubTypeInitSession)
	queryParams.Add("SessionID", sessionID)

	return "Command&" + queryParams.Encode()
}
