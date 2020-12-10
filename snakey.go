// Package snakey provides methods to convert text names to snake_case.
package snakey

import (
	"regexp"
	"strings"
)

// TextToSnake converts a given string to snake_case removing any spaces and non-alphanumeric characters
func TextToSnake(s string) string {
	var result string
	var words []string
	words = strings.Split(s, " ")
	for k, word := range words {
		if k > 0 {
			result += "_"
		}

		result += replaceChars(word)
	}

	return result
}

func replaceChars(s string) string {
	s = strings.ReplaceAll(s, " - ", "")
	s = strings.ReplaceAll(s, "-", "_")
	reg, err := regexp.Compile("[^A-Za-z0-9_]+")
	if err != nil {
		return s
	}
	return strings.ToLower(strings.TrimSpace(reg.ReplaceAllString(s, "")))
}
