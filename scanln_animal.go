package main

import (
	"fmt"
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

	for {
		fmt.Print("\nQuery: ")

		var animalName, info string
		_, err := fmt.Scanln(&animalName, &info)
		if err != nil {
			fmt.Println("Invalid Input. Use the format: <animal> <information>")
			fmt.Scanln()
			continue
		}

		// Handle "exit" command
		if strings.ToLower(animalName) == "exit" {
			fmt.Println("Thank you")
			break
		}

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

// In this code ScanLn is making problem when only exit is input to exit the code. Scanln expects two inputs.
//Since only one input is provided, it returns an error, leading to the "Invalid Input" message.
// To resolve this we will have to use bufio:- read the entire user input as a single string and then split it into components.
