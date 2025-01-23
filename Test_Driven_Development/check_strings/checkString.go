package main

import (
	"fmt"
	"strings"
	"unicode"
)

func checkString(input string) string {
	input = strings.TrimSpace(input)

	if len(input) == 0 || len(input) < 3 {
		return "Not Found!"
	}

	ifoundStart, nfoundEnd, afound := false, false, false
	n := len(input)

	for i, char := range input {
		if i == 0 && unicode.ToLower(char) == 'i' {
			ifoundStart = true
		}

		if i == n-1 && unicode.ToLower(char) == 'n' {
			nfoundEnd = true
		}

		if unicode.ToLower(char) == 'a' {
			afound = true
		}

		if ifoundStart && nfoundEnd && afound {
			return "Found!"
		}
	}
	return "Not Found!"
}

func main() {
	fmt.Println("Enter a String: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(checkString(input))
}
