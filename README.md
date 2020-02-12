# JWT Cracker written in Golang
This is a simple tool used to bruteforce HMAC secret keys in JWT tokens.
It heavly relies on this JWT library: `github.com/dgrijalva/jwt-go`

This tool supports both wordlist and bruteforce based attacks.

## Installation
```
go install github.com/ramosisw/go-jwt-cracker
```
You'll find the binary under the `bin` directory.

## Usage
This is a CLI tool, this means that it is meant to be used from your shell/terminal.

```txt
go-jwt-cracker is a simple tool used to bruteforce HMAC secret keys in JWT tokens

Usage:

	go-jwt-cracker [arguments]

The arguments are:
	--token <token>       The token you want to crack
	--wordlist <file>     The file for wordlist attack
	--charset <charset>   Specify the charset to use in the bruteforce attack
	--max <number>        The lower limit of the string's length for the brute force attack
	--min <number>        The upper limit of the string's length for the brute force attack
	--brute               Start the brute force attack

```

## Examples
In this examples I'll use this token: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.uJ0gEbC69N61cc2wHHl5vnC6uZMyLYUSwwQjwbYguqE`

that has been signed with the secret = `qwertyui`.

### Example 1 - Brute force
```
go-jwt-cracker --token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.uJ0gEbC69N61cc2wHHl5vnC6uZMyLYUSwwQjwbYguqE --brute
```

### Example 2 - Wordlist mode
```
go-jwt-cracker --token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.uJ0gEbC69N61cc2wHHl5vnC6uZMyLYUSwwQjwbYguqE --wordlist dict.txt[+] Valid secret found: secret
```