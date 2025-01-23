package main

import (
	"regexp"
	"strings"
)

func checkStringRegex(input string) string {
	input = strings.TrimSpace(input)
	pattern := `(?i)^i.*a.*n$`
	re := regexp.MustCompile(pattern)
	if re.MatchString(input) {
		return "Found!"
	}
	return "Not Found!"
}
