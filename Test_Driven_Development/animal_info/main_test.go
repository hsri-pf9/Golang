package main

import (
	"testing"
	// "errors"
)

// Test parseInput function
func TestParseInput(t *testing.T) {
	tests := []struct {
		input      string
		expected   []string
		expectError bool
	}{
		{"cow eat", []string{"cow", "eat"}, false},
		{"bird move", []string{"bird", "move"}, false},
		{"snake speak", []string{"snake", "speak"}, false},
		{"hello", nil, true}, // Invalid input (only one word)
		{"snake", nil, true}, // Missing information
		{"", nil, true}, // Empty input
		{" ", nil, true}, // Whitespace input
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := parseInput(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("parseInput(%q) error = %v, expectError %v", tt.input, err, tt.expectError)
			}
			if !tt.expectError && len(result) != 2 {
				t.Errorf("parseInput(%q) = %v, expected two fields", tt.input, result)
			}
		})
	}
}

// Test getAnimalInfo function
func TestGetAnimalInfo(t *testing.T) {
	animals := map[string]map[string]string{
		"cow": {"eat": "grass", "move": "walk", "speak": "moo"},
		"bird": {"eat": "worms", "move": "fly", "speak": "peep"},
		"snake": {"eat": "mice", "move": "slither", "speak": "hsss"},
	}

	tests := []struct {
		animal     string
		info       string
		expected   string
		expectError bool
	}{
		{"cow", "eat", "grass", false},
		{"bird", "move", "fly", false},
		{"snake", "speak", "hsss", false},
		{"cow", "sleep", "", true}, // Information doesn't exist
		{"cat", "eat", "", true}, // Animal doesn't exist
	}

	for _, tt := range tests {
		t.Run(tt.animal+" "+tt.info, func(t *testing.T) {
			result, err := getAnimalInfo(animals, tt.animal, tt.info)
			if (err != nil) != tt.expectError {
				t.Errorf("getAnimalInfo(%q, %q) error = %v, expectError %v", tt.animal, tt.info, err, tt.expectError)
			}
			if result != tt.expected {
				t.Errorf("getAnimalInfo(%q, %q) = %q, expected %q", tt.animal, tt.info, result, tt.expected)
			}
		})
	}
}