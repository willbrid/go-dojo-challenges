package wordhelper

import (
	"strings"
	"unicode"
)

func CountWordFrequency(text string) map[string]int {
	if len(text) > 10000 {
		return nil
	}

	var (
		result     map[string]int
		parsedText string
	)

	parsedText = strings.ToLower(text)
	replacer := strings.NewReplacer("\t", " ", "\n", " ", "-", " ")
	parsedText = replacer.Replace(parsedText)
	mapper := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != ' ' {
			return -1
		}

		return r
	}
	parsedText = strings.Map(mapper, parsedText)
	words := strings.Fields(parsedText)

	result = map[string]int{}
	for _, word := range words {
		result[word]++
	}

	return result
}
