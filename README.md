# Division1 tech challenge

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
    
    d1req -u nico -p myAwesomePassword
```
