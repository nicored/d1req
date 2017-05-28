package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"authentication"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	app := cli.NewApp()

	app.Name = "Div1 Authentication"
	app.Usage = "Check user credentials"
	app.UsageText = "auth [-u username] [-p password]"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "username, u",
			Value: "",
			Usage: "`Username` of the account",
		},
		cli.StringFlag{
			Name:  "password, p",
			Value: "",
			Usage: "`Password` of the account [Not recommended, use stdin instead]",
		},
	}

	app.Action = func(c *cli.Context) error {
		username, err := usernameHandler(c)
		if err != nil {
			return err
		}

		password, err := passwordHandler(c)
		if err != nil {
			return err
		}

		auth := authentication.Auth{
			BaseUrl: "test.div1.io",
			SSL:     true,
			Path:    "/test/1/exec",
		}

		_, err = auth.Authenticate(username, password)
		if err != nil {
			fmt.Println("\n" + err.Error())
			return err
		}

		fmt.Println(fmt.Sprintf("Hi %s, you are successfully authenticated", username))
		return nil
	}

	app.Run(os.Args)
}

func usernameHandler(c *cli.Context) (string, error) {
	var err error
	username := ""

	// Check that option is used
	if c.IsSet("username") {
		username = strings.TrimSpace(c.String("username"))
	}

	// If username is empty, we get it from stdin
	if username == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Username: ")

		username, err = reader.ReadString('\n')
		if err != nil {
			msg := "An error occurred reading your username"
			fmt.Println(msg)
			return "", errors.Wrap(err, msg)
		}

		if username == "" {
			msg := "Username cannot be empty"
			fmt.Println(msg)
			return "", errors.Wrap(err, msg)
		}

		username = username[:len(username)-1]
	}

	return username, nil
}

func passwordHandler(c *cli.Context) (string, error) {
	var err error
	password := ""

	if c.IsSet("password") {
		fmt.Println("WARNING: It is not safe, and therefore not recommended to enter the password in command line arguments. Use stdin instead.")
		password = c.String("password")
	}

	// If password is empty, we get it from stdin
	if password == "" {
		var bytePassword []byte

		fmt.Print("Password: ")

		// We want to hide the password in terminal
		bytePassword, err = terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			msg := "An error occurred reading your password"
			fmt.Println(msg)
			return "", errors.Wrap(err, msg)
		}
		password = string(bytePassword)

		if password == "" {
			msg := "Password cannot be empty"
			fmt.Println(msg)
			return "", errors.Wrap(err, msg)
		}
	}

	return password, nil
}
