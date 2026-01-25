package main

import (
	"fmt"

	"calculator/sum"
)

func main() {
	var a, b int
	fmt.Println("Veuillez entrer deux nombres")
	_, err := fmt.Scanf("%d", "%d", &a, &b)
	if err != nil {
		fmt.Println("Error reading input :", err)
		return
	}

	result := sum.Sum(a, b)
	fmt.Printf("%d + %d = %d", a, b, result)
}
