# Division1 tech challenge

[![Go Report Card](https://goreportcard.com/badge/github.com/nicored/d1req)](https://goreportcard.com/report/github.com/nicored/d1req) [![Build Status](https://travis-ci.org/nicored/d1req.svg)](https://travis-ci.org/nicored/d1req) [![Coverage Status](https://coveralls.io/repos/github/nicored/d1req/badge.svg?branch=master)](https://coveralls.io/github/nicored/d1req?branch=master)

##### GoDocs: 
Authentication: [![GoDoc](https://godoc.org/github.com/nicored/d1req/src/authentication?status.svg)](https://godoc.org/github.com/nicored/d1req/src/authentication) 

Encryption: [![GoDoc](https://godoc.org/github.com/nicored/d1req/src/encryption?status.svg)](https://godoc.org/github.com/nicored/d1req/src/encryption)

## Overview

The challenge consisted in reverse-engineering minified code from a
a javascript client and implement the functionalities in golang.

## Issues encountered

The javascript client sends the username in lower-case, so I assumed
that the username would be case-insensitive, and guess what... I was wrong.

## Installation

To install the binary, run the following command. This will create d1req exec
file in $GOBIN:

```shell
    go get github.com/nicored/d1req
```

## Usage

```shell
    # You can check out how to use the cli
    d1req help
```

```shell
    # --username (-u) and --password (-p) are optional
    # If they're not provided, or empty, you will be asked to enter them in stdin
    # The program will warn you that typing your pswd as a cmd line argument is bad
    # Also, don't you worry, passwords typed in standard input won't show up
    
    $ d1req -u nico -p myAwesomePassword
    WARNING: It is not safe, and therefore not recommended to enter the password in command line arguments. Use stdin instead.
    Hi nico, you are successfully authenticated
    
    $ d1req -u nico
    Password: 
    Hi nico, you are successfully authenticated
    
    $ d1req
    Username: nico
    Password: 
    Hi nico, you are successfully authenticated
    
    $ d1req -u wrongusername
    Password:
    Authentication failed.: Invalid credentials.
```

## Performance

### Xor

Performance of Xor is quite poor. Converting and parsing data to different types
and formats is very expensive, as well as string operations. 

I did not have much time to look much into alternatives, but I believe that
more efficient ways to do this exist. I'd give it another go by operating with bytes only
where possible, which I expect would slightly improve the performance.

        Average   Total over 2.45s
        90ms      2.45s (flat, cum) 99.19% of Total
         .          .     27:// num is the number used to perform xor comparisons with the input
         .          .     28:// The output of Xor is a string representing the result of the operation in Hexadecimal format
         .          .     29:// in UPPERCASE
         .          .     30:func Xor(input string, num int64) (string, error) {
         .          .     31:	// Convert num to binary format
         .      250ms     32:	numBin := fmt.Sprintf("%032s", strconv.FormatInt(num, 2))
         .          .     33:
         .          .     34:	startPos := len(numBin)
         .          .     35:	retVal := ""
         .          .     36:	for _, char := range input {
         .          .     37:		if startPos == 0 {
         .          .     38:			startPos = len(numBin) - 8
         .          .     39:		} else {
         .          .     40:			startPos = startPos - 8
         .          .     41:		}
         .          .     42:
         .          .     43:		// Convert 1 byte at a time to int64 format
      20ms      320ms     44:		comp, err := strconv.ParseInt(numBin[startPos:startPos+8], 2, 64)
         .          .     45:		if err != nil {
         .          .     46:			return "", errors.Wrap(err, "Error converting byte to int.")
         .          .     47:		}
         .          .     48:
         .          .     49:		// Xor operation on char code and current byte in num
         .          .     50:		xorInt := int64(char) ^ comp
         .          .     51:
         .          .     52:		// Convert xor result to hexadecimal formal
         .          .     53:		// and append '0' if xorHex <= F
      20ms      570ms     54:		xorHex := strconv.FormatInt(xorInt, 16)
      10ms       10ms     55:		if len(xorHex) == 1 {
      10ms       70ms     56:			xorHex = "0" + xorHex
         .          .     57:		}
         .          .     58:
         .          .     59:		// Append to the chain
      20ms      820ms     60:		retVal += xorHex
         .          .     61:	}
         .          .     62:
      10ms      410ms     63:	return strings.ToUpper(retVal), nil
         .          .     64:}