package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

func printHelp() {
	fmt.Println(`go-jwt-cracker is a simple tool used to bruteforce HMAC secret keys in JWT tokens

Usage:

	go-jwt-cracker [arguments]

The arguments are:
	--token <token>       The token you want to crack
	--wordlist <file>     The file for wordlist attack
	--charset <charset>   Specify the charset to use in the bruteforce attack
	--max <number>        The lower limit of the string's length for the brute force attack
	--min <number>        The upper limit of the string's length for the brute force attack
	--brute               Start the brute force attack
 `)
}

func main() {

	var tokenString = flag.String("token", "", "The JWT token you want to crack")
	var wordlist = flag.String("wordlist", "", "The wordlist you want to use")
	var brute = flag.Bool("brute", false, "Use bruteforce mode")
	var charset = flag.String("charset", "qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM", "The charset to use during bruteforce")
	var minChar = flag.Int("min", 7, "Min chars in token")
	var maxChar = flag.Int("max", 12, "Max chars in token")

	flag.Parse()

	if string(*tokenString) == "" {
		printHelp()
		return
	}

	fmt.Println("Cracking...")

	if *wordlist != "" && *brute == false {
		wordlistFile, _ := os.Open(*wordlist)
		defer wordlistFile.Close()

		scanner := bufio.NewScanner(wordlistFile)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			currentToken := scanner.Text()
			if validateToken(*tokenString, currentToken) {
				fmt.Println("[+] Valid secret found: " + currentToken)
				return
			}
		}
	}

	if *brute == true {
		for lenChar := *minChar; lenChar <= *maxChar; {
			fmt.Printf("Trying with lentgth: %v \n", lenChar)
			for combination := range generateCombinations(*charset, lenChar) {
				if validateToken(*tokenString, combination) {
					fmt.Println("[+] Valid secret found: " + combination)
					return
				}
			}
		}
	}

}

// I found this function in a stackoverflow answer that I'm not finding anymore
// when I'll find it I'll give appropriate credits to the author
func generateCombinations(alphabet string, length int) <-chan string {
	c := make(chan string)
	go func(c chan string) {
		defer close(c)

		addLetter(c, "", alphabet, length)
	}(c)
	return c
}

func addLetter(c chan string, combo string, alphabet string, length int) {
	if length <= 0 {
		return
	}

	var newCombo string
	for _, ch := range alphabet {
		newCombo = combo + string(ch)
		c <- newCombo
		addLetter(c, newCombo, alphabet, length-1)
	}
}

func validateToken(tokenString string, secret string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		hmacSampleSecret := []byte(secret)
		return hmacSampleSecret, nil
	})
	return token.Valid
}
