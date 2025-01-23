package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	animals := map[string]map[string]string{
		"cow":   {"eat": "grass", "move": "walk", "speak": "moo"},
		"bird":  {"eat": "worms", "move": "fly", "speak": "peep"},
		"snake": {"eat": "mice", "move": "slither", "speak": "hsss"},
	}

	fmt.Println("Enter queries in the format: <animal> <information>")
	fmt.Println("Example: cow eat")
	fmt.Println("To exit, type 'exit'.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nQuery: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}

		input = strings.TrimSpace(input)

		// Handle "exit" command
		if strings.ToLower(input) == "exit" {
			fmt.Println("Thank you")
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid Input. Use the format: <animal> <information>")
			continue
		}

		animalName, info := parts[0], parts[1]

		animal, exists := animals[animalName]
		if !exists {
			fmt.Printf("Animal '%s' not found.\n", animalName)
			continue
		}

		value, infoExists := animal[info]
		if !infoExists {
			fmt.Printf("Information '%s' not found for animal '%s'.\n", info, animalName)
			continue
		}

		fmt.Printf("The %s %s: %s\n", animalName, info, value)
	}
}

// Now the code is working perfectly fine.