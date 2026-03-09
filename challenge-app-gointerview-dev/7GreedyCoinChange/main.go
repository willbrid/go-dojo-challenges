package main

import (
	"fmt"

	"greedycoinchange/coin"
)

func main() {
	denominations := []int{1, 5, 10, 25, 50}
	amounts := []int{87, 42, 99, 33, 7}

	for _, amount := range amounts {
		minCoins := coin.MinCoins(amount, denominations)
		coinCombo := coin.CoinCombination(amount, denominations)

		fmt.Printf("Amount: %d cents\n", amount)
		fmt.Printf("Minimum coins needed: %d\n", minCoins)
		fmt.Printf("Coin combination: %v\n", coinCombo)
		fmt.Println("---------------------------")
	}
}
