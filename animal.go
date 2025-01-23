package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.sound
}

func main() {
	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter queries in the format: <animal> <information>")
	fmt.Println("Example: cow eat")
	fmt.Println("To exit, type 'exit'.")

	for {
		fmt.Print("\nQuery: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Thankyou")
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input")
			continue
		}

		animalName, info := parts[0], parts[1]

		animal, exists := animals[animalName]
		if !exists {
			fmt.Printf("Animal '%s' not found. \n", animalName)
			continue
		}
		switch info {
		case "eat":
			fmt.Printf("The %s eats: %s\n", animalName, animal.Eat())
		case "move":
			fmt.Printf("The %s moves: %s\n", animalName, animal.Move())
		case "speak":
			fmt.Printf("The %s speaks: %s\n", animalName, animal.Speak())
		default:
			fmt.Printf("Information '%s' not recognized", info)
		}
	}
}

// In this code i have used struct and map
