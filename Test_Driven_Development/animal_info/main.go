package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
)

func parseInput(input string) ([]string, error) {
	parts := strings.Fields(input)
	if len(parts) != 2 {
		return nil, errors.New("invalid input")
	}
	return parts, nil
}

func getAnimalInfo(animals map[string]map[string]string, animal string, info string) (string, error) {
	animalData, exists := animals[animal]
	if !exists {
		return "", fmt.Errorf("animal '%s' not found", animal)
	}

	value, infoExists := animalData[info]
	if !infoExists {
		return "", fmt.Errorf("information '%s' not found for animal '%s'", info, animal)
	}

	return value, nil
}

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

		parts, err := parseInput(input)
		if err != nil {
			fmt.Println("Invalid Input. Use the format: <animal> <information>")
			continue
		}

		animalName, info := parts[0], parts[1]

		value, err := getAnimalInfo(animals, animalName, info)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("The %s %s: %s\n", animalName, info, value)
	}
}
