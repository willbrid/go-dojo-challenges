package reversestring_test

import (
	"stringparser/reversestring"
	"testing"
)

type testCase struct {
	name     string
	input    string
	expected string
}

func TestReverseString_ReturnCorrect(t *testing.T) {
	tests := []testCase{
		{"Simple word", "hello", "olleh"},
		{"Sentence with spaces", "Go is fun!", "!nuf si oG"},
		{"Empty string", "", ""},
		{"Palindrome", "madam", "madam"},
		{"Special characters", "12345!@#$%", "%$#@!54321"},
		{"Mixed case", "GoLang", "gnaLoG"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(subT *testing.T) {
			result := reversestring.ReverseString(tt.input)
			if result != tt.expected {
				subT.Errorf("ReverseString(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
