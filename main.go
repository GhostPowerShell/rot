package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
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
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s plaintext key\n", os.Args[0])
		return
	}

	plaintext := os.Args[1]
	key, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid key:", os.Args[2])
		return
	}

	ciphertext := rotN(plaintext, key)

	var output bytes.Buffer
	output.WriteString(whiteBoldColor)
	output.WriteString("Your text: ")
	output.WriteString(redColor)
	output.WriteString(ciphertext)
	output.WriteString("\n")
	output.WriteString(resetColor)

	os.Stdout.Write(output.Bytes())
}
