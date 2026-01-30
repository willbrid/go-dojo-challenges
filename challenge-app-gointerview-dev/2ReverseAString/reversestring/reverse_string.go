package reversestring

func ReverseString(s string) string {
	var result string = ""
	length := len(s)

	for i := length - 1; i >= 0; i-- {
		result += "" + string(s[i])
	}

	return result
}
