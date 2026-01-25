package sum

func Sum(a, b int) int {
	if a < -1000000000 && b > 1000000000 {
		panic("a and b must be between -1e9 and 1e9")
	}

	return a + b
}
