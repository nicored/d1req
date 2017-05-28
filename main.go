package main

import (
	"authentication"
)

func main() {
	auth := authentication.Auth{
		BaseUrl: "test.div1.io",
		SSL:     true,
		Path:    "/test/1/exec",
	}

	username := "ClubSandwich"
	password := "T$(LCy gbhn3wb798"

	_, err := auth.Authenticate(username, password)
	if err != nil {
		panic(err)
	}
}
