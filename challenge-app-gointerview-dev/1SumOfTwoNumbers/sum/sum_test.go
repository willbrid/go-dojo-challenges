package sum_test

import (
	"testing"

	"calculator/sum"
)

type TestCase struct {
	name     string
	a, b     int
	expected int
}

func TestSum_ReturnCorrectResponse(t *testing.T) {
	tests := []TestCase{
		{"positive numbers", 2, 3, 5},
		{"zero values", 0, 0, 0},
		{"negative numbers", -2, -3, -5},
		{"mix numbers 1", -2, 3, 1},
		{"mix numbers 2", 2, -3, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(subT *testing.T) {
			result := sum.Sum(tt.a, tt.b)
			if result != tt.expected {
				subT.Errorf("Sum(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSum_PanicWhenConditionIsFailed(t *testing.T) {
	var a, b int = -10000000000, 10000000000

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Sum(%d, %d) should panic", a, b)
		}
	}()

	sum.Sum(a, b)
}
