package main

import (
	"fmt"

	"binarysearch/search"
)

func main() {
	// Example sorted array for testing
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// Test binary search
	target := 7
	index := search.BinarySearch(arr, target)
	fmt.Printf("BinarySearch: %d found at index %d\n", target, index)

	// Test recursive binary search
	recursiveIndex := search.BinarySearchRecursive(arr, target, 0, len(arr)-1)
	fmt.Printf("BinarySearchRecursive: %d found at index %d\n", target, recursiveIndex)

	// Test find insert position
	insertTarget := 8
	insertPos := search.FindInsertPosition(arr, insertTarget)
	fmt.Printf("FindInsertPosition: %d should be inserted at index %d\n", insertTarget, insertPos)
}
