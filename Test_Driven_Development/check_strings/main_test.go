package main

import (
	"testing"
	"time"
)

func TestCheckString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"ian", "Found!"},
		{"Ian", "Found!"},
		{"iuiygaygn", "Found!"},
		{"I d skd a efju N", "Found!"},
		{"ihhhhhn", "Not Found!"},
		{"ina", "Not Found!"},
		{"xian", "Not Found!"},
	}

	// Test checkString
	t.Run("Normal Approach", func(t *testing.T) {
		start := time.Now()

		for _, tt := range tests {
			t.Run(tt.input, func(t *testing.T) {
				result := checkString(tt.input)
				if result != tt.expected {
					t.Errorf("checkStringNormal(%q) = %q; want %q", tt.input, result, tt.expected)
				}
			})
		}

		duration := time.Since(start)
		t.Logf("Total time taken for Normal Approach: %v", duration)
	})

	// Test checkStringRegex
	t.Run("Regex Approach", func(t *testing.T) {
		start := time.Now()

		for _, tt := range tests {
			t.Run(tt.input, func(t *testing.T) {
				result := checkStringRegex(tt.input)
				if result != tt.expected {
					t.Errorf("checkStringRegex(%q) = %q; want %q", tt.input, result, tt.expected)
				}
			})
		}

		duration := time.Since(start)
		t.Logf("Total time taken for Regex Approach: %v", duration)
	})
}
