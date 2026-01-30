package main

import (
	"bufio"
	"fmt"
	"os"

	"stringparser/reversestring"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		output := reversestring.ReverseString(input)
		fmt.Println(output)
	}
}
