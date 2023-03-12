package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

const (
	whiteBoldColor = "\033[1;37m"
	redColor       = "\033[1;31m"
	resetColor     = "\033[0m"
)

func rotN(input string, key int) string {
	var output bytes.Buffer
	for _, i := range input {
		if i >= 'a' && i <= 'z' {
			i = 'a' + (i-'a'+rune(key))%26
		} else if i >= 'A' && i <= 'Z' {
			i = 'A' + (i-'A'+rune(key))%26
		}
		output.WriteRune(i)
	}
	return output.String()
}

func main() {
	var plaintext string
	var key int
	flag.StringVar(&plaintext, "t", "", "the plaintext to encrypt")
	flag.IntVar(&key, "k", 0, "the encryption key")
	flag.Parse()

	if plaintext == "" || key == 0 {
		fmt.Fprintln(os.Stderr, "Usage: rot -t plaintext -k key")
		os.Exit(1)
	}
	if key < 0 {
		fmt.Fprintln(os.Stderr, "Invalid key:", key)
		os.Exit(1)
	}

	ciphertext := rotN(plaintext, key)

	output := fmt.Sprintf("%sYour text: %s%s\n%s", whiteBoldColor, redColor, ciphertext, resetColor)

	if _, err := os.Stdout.WriteString(output); err != nil {
		fmt.Fprintln(os.Stderr, "Error writing output:", err)
		os.Exit(1)
	}

}
