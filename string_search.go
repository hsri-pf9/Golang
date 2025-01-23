package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println("Enter a string: ")
	var input string
	fmt.Scanln(&input)

	lowerInput := strings.ToLower(input)

	if strings.HasPrefix(lowerInput, "i") && strings.HasSuffix(lowerInput, "n") && strings.Contains(lowerInput, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}