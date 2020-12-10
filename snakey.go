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
	var lastPos int
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		words = append(words, replaceChars(s[lastPos:i]))
		lastPos = i
	}

	// append the last word
	if s[lastPos:] != "" {
		words = append(words, replaceChars(s[lastPos:]))
	}

	for k, word := range words {
		if k > 0 {
			result += "_"
		}

		result += strings.ReplaceAll(strings.ToLower(strings.TrimSpace(word)), "-", "_")
	}

	return result
}

func replaceChars(s string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9`-`]+")
	if err != nil {
		return s
	}
	return reg.ReplaceAllString(s, "")
}
